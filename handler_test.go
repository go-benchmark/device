package device

import (
	"context"
	"fmt"
	"testing"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-benchmark/config"
	"github.com/go-benchmark/service"
	"github.com/stretchr/testify/assert"
)

func newDevice(t *testing.T) *Device {
	vu := 1
	opts, err := config.ConfigureOptions(vu)
	assert.Nil(t, err)
	d, err := NewDevice(opts, 1,nil)
	assert.Nil(t, err)

	return d
}

type nullMqttClient struct {
	pub struct {
		topic string
		qos   byte
		data  []byte
	}
}

func (n *nullMqttClient) Connect(ctx context.Context) error {
	return nil
}

func (n *nullMqttClient) Publish(ctx context.Context, topic string, qos byte, data []byte) error {
	n.pub = struct {
		topic string
		qos   byte
		data  []byte
	}{
		topic: topic,
		qos:   qos,
		data:  data,
	}

	return nil
}

func (n *nullMqttClient) Subscribe(ctx context.Context, topic string, qos byte, callback paho.MessageHandler) error {
	return nil
}

func (n *nullMqttClient) getLastPublish() (string, byte, []byte) {
	return n.pub.topic, n.pub.qos, n.pub.data
}

func newNullMqttClient() *nullMqttClient {
	return &nullMqttClient{}
}

func TestHandleConfigPayload(t *testing.T) {
	ctx := context.TODO()
	payload := `{"engineType":"softSecurity","serviceId":"123","cmd":"start"}`
	d := newDevice(t)
	mc := newNullMqttClient()

	err := d.handleConfigPayload(ctx, mc, []byte(payload))
	assert.Nil(t, err)

	s, ok := d.services["123"]
	assert.True(t, ok)
	assert.Equal(t, service.Service{
		ServiceID:  "123",
		EngineType: "softSecurity",
		OI:         service.OperationInfo{
			// Run:    false,
			// Saving: false,
			// Bots:   []string{},
		},
	}, s)
}

func TestResHistory(t *testing.T) {
	ctx := context.TODO()
	payload := `{"engineType":"softSecurity","serviceId":"123","cmd":"start"}`
	d := newDevice(t)
	mc := newNullMqttClient()

	err := d.handleConfigPayload(ctx, mc, []byte(payload))
	assert.Nil(t, err)

	err = d.resHistory(ctx, mc, "123")
	assert.Nil(t, err)

	topic, qos, _ := mc.getLastPublish()

	assert.Equal(t, fmt.Sprintf("org/vz-loadtest/%s/opState", d.ID), topic)
	assert.Equal(t, byte(1), qos)
}

func TestDevice_handleConfigPayload(t *testing.T) {
	type args struct {
		ctx     context.Context
		mc      mqtter
		payload []byte
	}
	tests := []struct {
		name    string
		d       *Device
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.handleConfigPayload(tt.args.ctx, tt.args.mc, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Device.handleConfigPayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_sub(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		d       *Device
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.sub(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Device.sub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_pub(t *testing.T) {
	type args struct {
		ctx     context.Context
		topic   string
		qos     byte
		payload []byte
	}
	tests := []struct {
		name    string
		d       *Device
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.pub(tt.args.ctx, tt.args.topic, tt.args.qos, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Device.pub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
