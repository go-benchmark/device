package device

import (
	"encoding/json"

	"github.com/bxcodec/faker/v3"
)

type property struct {
	IP          string `faker:"ipv4" json:"ip"`
	Description string `faker:"sentence" json:"description"`
}

type node struct {
	ID         string   `faker:"mac_address" json:"id"`
	Properties property `json:"properties"`
}
type link struct {
	Source string  `faker:"mac_address" json:"source"`
	Target string  `faker:"mac_address" json:"target"`
	Cost   float64 `json:"cost"`
}
type status struct {
	Type     string `faker:"word" json:"type"`
	Protocol string `faker:"word" json:"protocol"`
	Version  string `json:"version"`
	Metric   string `json:"metric"`
	Nodes    []node `json:"nodes"`
	Links    []link `json:"links"`
}

// FakeStatus returns fake reporting status
// support  "type": "NetworkGraph"
func FakeStatus() (s []byte, err error) {
	ss := status{}

	faker.SetRandomMapAndSliceSize(10)

	if err = faker.FakeData(&ss); err != nil {
		return
	}

	ss.Type = "NetworkGraph"

	s, err = json.Marshal(ss)
	if err != nil {
		return
	}

	return
}
