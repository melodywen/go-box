package support

import "testing"

func TestStrFun_Is(t *testing.T) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试 空h支付串",
			args: args{
				pattern: "",
				value:   "aaa",
			},
			want: false,
		}, {
			name: "测试 正则匹配",
			args: args{
				pattern: "m*y",
				value:   "melody",
			},
			want: true,
		}, {
			name: "测试 不是边界",
			args: args{
				pattern: "e*y",
				value:   "melody",
			},
			want: false,
		}, {
			name: "测试 正则匹配",
			args: args{
				pattern: "m*d",
				value:   "melody",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := StrFun{}
			if got := st.Is(tt.args.pattern, tt.args.value); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
