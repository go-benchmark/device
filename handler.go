package device

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-benchmark/service"
)

type engineMsg struct {
	Realtime       bool `json:"realtime"`
	RealtimeLength int  `json:"realtimeLength"`
}

type cfgPayload struct {
	service.Service
	Cmd       string
	EngineMsg engineMsg
}

func (d *Device) handleConfigPayload(ctx context.Context, mc mqtter, payload []byte) (err error) {
	cp := cfgPayload{}
	if err = json.Unmarshal(payload, &cp); err != nil {
		return
	}

	// https://github.com/originwc/ow-verik-doc/blob/master/reference-system/device/mqtt-spec.md#413-ss-sending-startstop-service-request
	// Start/Stop Service Request
	// NOTE: only support start
	// stop is not supported yet
	if cp.Cmd == "start" {
		err = d.cfgStartService(ctx, mc, cp)
		return
	}

	// handle heartbeat command from user
	if cp.EngineMsg.Realtime {
		s, ok := d.services[cp.ServiceID]
		if !ok {
			return ErrServiceNotFound
		}
		// send hearbeat to heartbeat channel
		go func(ctx context.Context) {
			err = s.RealtimeConfig(ctx)
			return
		}(ctx)
	}
	return
}

func (d *Device) handleConfig(ctx context.Context) func(paho.Client, paho.Message) {
	return func(client paho.Client, msg paho.Message) {
		log.Printf("[%d] config handler receiving: %s\n", d.vu, string(msg.Payload()))

		if err := d.handleConfigPayload(ctx, &d.mc, msg.Payload()); err != nil {
			log.Printf("failed handleConfigPayload: %v\n", err)
		}
	}
}

func (d *Device) sub(ctx context.Context) (err error) {
	topic := fmt.Sprintf("%s/%s/config", d.mqttPrefix, d.ID)
	err = d.mc.Subscribe(ctx, topic, 1, d.handleConfig(ctx))

	return
}
func (d *Device) pub(ctx context.Context, topic string, qos byte, payload []byte) (err error) {
	if topic == "" {
		// use default topic
		topic = d.opStateTopic()
	}
	log.Printf("{dev_pub}-id: %s |topic: %s |payload: %v", d.ID, topic, string(payload))
	return d.mc.Publish(ctx, topic, qos, payload)
}