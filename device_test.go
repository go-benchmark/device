package device

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-benchmark/config"
)

func TestNewDevice(t *testing.T) {
	type args struct {
		opts *config.Options
		vu   int
	}
	tests := []struct {
		name    string
		args    args
		want    *Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDevice(tt.args.opts, tt.args.vu,nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_setOpts(t *testing.T) {
	type args struct {
		opts *config.Options
		vu   int
	}
	tests := []struct {
		name string
		d    *Device
		args args
		want *Device
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.setOpts(tt.args.opts, tt.args.vu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Device.setOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetMqttAuth(t *testing.T) {
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
			if err := tt.d.GetMqttAuth(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Device.GetMqttAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_ConnectMqtt(t *testing.T) {
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
			if err := tt.d.ConnectMqtt(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Device.ConnectMqtt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_Run(t *testing.T) {
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
			if err := tt.d.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Device.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_Finish(t *testing.T) {
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
			if err := tt.d.Finish(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Device.Finish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_opStateTopic(t *testing.T) {
	tests := []struct {
		name string
		d    *Device
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.opStateTopic(); got != tt.want {
				t.Errorf("Device.opStateTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}
