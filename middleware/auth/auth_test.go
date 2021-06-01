package auth

import (
	"context"
	"reflect"
	"testing"
)

func TestGateWayAuthFunc(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    context.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GateWayAuthFunc(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GateWayAuthFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GateWayAuthFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
