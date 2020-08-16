package device

import (
	"context"
	"reflect"
	"testing"
)

func TestDevice_activeBots(t *testing.T) {
	type args struct {
		ctx  context.Context
		bots []string
	}
	tests := []struct {
		name    string
		d       *Device
		args    args
		wantAb  ActiveBot
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAb, err := tt.d.activeBots(tt.args.ctx, tt.args.bots)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.activeBots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAb, tt.wantAb) {
				t.Errorf("Device.activeBots() = %v, want %v", gotAb, tt.wantAb)
			}
		})
	}
}
