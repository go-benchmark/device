package device

import (
	"context"
	"encoding/json"

	"github.com/bxcodec/faker/v3"
)

type netConf struct {
	Timestamp int64   `faker:"unix_time" json:"timestamp"`
	Network   network `json:"network"`
}

type network struct {
	Origin string   `json:"origin"`
	Bots   []string `json:"bots"`
}

// networkconfig create a faking network config struct
// then publish to the MQTT broker
func (d *Device) networkconfig(ctx context.Context) (nc netConf, err error) {
	if err = faker.FakeData(&nc); err != nil {
		return
	}

	nc.Network.Origin = d.ID
	nc.Network.Bots = []string{faker.UUIDDigit(), faker.UUIDDigit()}

	ncStr, err := json.Marshal(nc)

	if err != nil {
		return
	}

	opStateTopic := d.opStateTopic()
	if err = d.pub(ctx, opStateTopic, 1, ncStr); err != nil {
		return
	}

	return
}
