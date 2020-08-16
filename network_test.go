package device

import (
	"context"
	"reflect"
	"testing"
)

func TestDevice_networkconfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		d       *Device
		args    args
		wantNc  netConf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNc, err := tt.d.networkconfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.networkconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNc, tt.wantNc) {
				t.Errorf("Device.networkconfig() = %v, want %v", gotNc, tt.wantNc)
			}
		})
	}
}
