package utils

import (
	"reflect"
	"testing"
)

func TestRunTimeCost(t *testing.T) {
	tests := []struct {
		name string
		want func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunTimeCost(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunTimeCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
