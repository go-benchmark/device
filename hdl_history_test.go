package device

import (
	"context"
	"testing"
)

func TestDevice_resHistory(t *testing.T) {
	type args struct {
		ctx       context.Context
		mc        mqtter
		serviceID string
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
			if err := tt.d.resHistory(tt.args.ctx, tt.args.mc, tt.args.serviceID); (err != nil) != tt.wantErr {
				t.Errorf("Device.resHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
