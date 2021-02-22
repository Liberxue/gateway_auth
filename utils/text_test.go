package utils

import "testing"

func TestIsChineseChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsPass",
			args: args{
				str: "are you ok ",
			},
			want: false,
		},
		{
			name: "IsPass",
			args: args{
				str: "are you ok 哈",
			},
			want: true,
		},
		{
			name: "IsPass",
			args: args{
				str: "are 哈 you ok 哈",
			},
			want: true,
		},
		{
			name: "IsPass",
			args: args{
				str: "你好吗",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseChar(tt.args.str); got != tt.want {
				t.Errorf("IsChineseChar() = %v,want %v args %v", got, tt.want, tt.args)
			}
		})
	}
}

func TestIsEnglishChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsPass",
			args: args{
				str: "are aaa",
			},
			want: true,
		},
		{
			name: "IsPass",
			args: args{
				str: "this's ok",
			},
			want: true,
		},
		{
			name: "IsPass",
			args: args{
				str: "this's ok!",
			},
			want: true,
		},
		{
			name: "IsPass",
			args: args{
				str: "this's ok.",
			},
			want: true,
		},
		{
			name: "test Failed",
			args: args{
				str: "are you ok 哈",
			},
			want: false,
		},
		{
			name: "test Failed",
			args: args{
				str: "点点滴滴",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnglishChar(tt.args.str); got != tt.want {
				t.Errorf("IsEnglishChar() = %v, want %v args %v", got, tt.want, tt.args)
			}
		})
	}
}
