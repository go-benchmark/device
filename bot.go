package device

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
)

type ActiveBot struct {
	Timestamp int64     `json:"timestamp"`
	BotInfo   []BotInfo `json:"bot_info"`
}
type BotInfo struct {
	Bot    string `json:"bot"`
	Hwaddr string `json:"hwaddr"`
	Chains int    `json:"chains"`
}

// trick to set device online, dont cate about bots
func (d *Device) activeBots(ctx context.Context, bots []string) (ab ActiveBot, err error) {
	if err = faker.FakeData(&ab); err != nil {
		return
	}

	for _, bot := range bots {
		bi := BotInfo{
			Bot:    bot,
			Hwaddr: faker.UUIDDigit(),
		}
		ab.BotInfo = append(ab.BotInfo, bi)
	}
	// timestamp require only unix timestamp format with 10 digits, cant use faker
	ab.Timestamp = time.Now().UTC().Unix()
	abStr, err := json.Marshal(ab)
	if err != nil {
		return
	}

	opStateTopic := fmt.Sprintf("%s/%s/opState", d.mqttPrefix, d.ID)
	if err = d.pub(ctx, opStateTopic, 1, abStr); err != nil {
		return
	}

	return
}
