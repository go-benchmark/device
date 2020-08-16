package device

import (
	"reflect"
	"testing"
)

func TestFakeStatus(t *testing.T) {
	tests := []struct {
		name    string
		wantS   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := FakeStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("FakeStatus() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
