package device

import (
	"context"
	"errors"
	"log"
)

func (d *Device) resHistory(ctx context.Context, mc mqtter, serviceID string) (err error) {
	s, ok := d.services[serviceID]
	if !ok {
		return errors.New("service not found")
	}

	hBytes, err := s.History()
	if err != nil {
		return
	}

	topic := d.opStateTopic()

	if err = d.pub(ctx, topic, 1, hBytes); err != nil {
		log.Printf("error publish mqtt history message: %v \n", err)
		return
	}

	return
}
