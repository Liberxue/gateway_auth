package utils

import "testing"

func Test_limit_Run(t *testing.T) {
	type fields struct {
		Num int
		C   chan struct{}
	}
	type args struct {
		f func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &limit{
				Num: tt.fields.Num,
				C:   tt.fields.C,
			}
			g.Run(tt.args.f)
		})
	}
}
