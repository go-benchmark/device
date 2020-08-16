package device

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-benchmark/service"
	"github.com/gobench-io/gobench/dis"
)

type configuration struct {
	Timestamp  int64              `json:"timestamp"`
	EngineType service.EngineType `json:"engineType"`
	ServiceID  string             `json:"serviceId"`
	Cmd        string             `json:"cmd"`
	Success    bool               `json:"success"`
}

// cfgStartService generates a start response, save the service, and begin
// history routine
func (d *Device) cfgStartService(ctx context.Context, mc mqtter, cp cfgPayload) (err error) {
	if err = d.resStartService(ctx, mc, cp); err != nil {
		return
	}
	var s *service.Service

	if _, ok := d.services[cp.ServiceID]; !ok {
		// get default realtime length from config
		realtimeLength := d.dc.realtimeLength
		if cp.EngineMsg.RealtimeLength > 0 {
			realtimeLength = cp.EngineMsg.RealtimeLength
		}
		s = service.NewService(&cp.Service, realtimeLength)
		if err = d.saveService(s); err != nil {
			return
		}
	} else {
		s = d.services[cp.ServiceID]
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dis.SleepRatePoisson(1.0 / float64(d.historyInterval))
				d.resHistory(ctx, mc, cp.ServiceID)
			}
		}
	}()
	go func(ctx context.Context) {
		if err = s.RealtimeHandler(ctx, d.dc.realtimeInterval, d.pub); err != nil {
			return
		}
	}(ctx)

	return
}
func (d *Device) saveService(s *service.Service) (err error) {
	d.services[s.ServiceID] = s
	return
}

func (d *Device) resStartService(ctx context.Context, mc mqtter, cfgPayload cfgPayload) (err error) {
	c := configuration{}
	c.Timestamp = time.Now().Unix()
	c.EngineType = cfgPayload.EngineType
	c.ServiceID = cfgPayload.ServiceID
	c.Cmd = "start"
	c.Success = true

	cBytes, err := json.Marshal(c)
	if err != nil {
		return
	}

	opStateTopic := d.opStateTopic()
	if err = d.pub(ctx, opStateTopic, 1, cBytes); err != nil {
		return
	}

	log.Printf("[%d] responded configuration:%s \n", d.vu, string(cBytes))

	return
}
