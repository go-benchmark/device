package device

import (
	"context"
	"testing"

	"github.com/go-benchmark/service"
)

func TestDevice_cfgStartService(t *testing.T) {
	type args struct {
		ctx context.Context
		mc  mqtter
		cp  cfgPayload
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
			if err := tt.d.cfgStartService(tt.args.ctx, tt.args.mc, tt.args.cp); (err != nil) != tt.wantErr {
				t.Errorf("Device.cfgStartService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_saveService(t *testing.T) {
	type args struct {
		s *service.Service
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
			if err := tt.d.saveService(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Device.saveService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDevice_resStartService(t *testing.T) {
	type args struct {
		ctx        context.Context
		mc         mqtter
		cfgPayload cfgPayload
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
			if err := tt.d.resStartService(tt.args.ctx, tt.args.mc, tt.args.cfgPayload); (err != nil) != tt.wantErr {
				t.Errorf("Device.resStartService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
