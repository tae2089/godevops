package redis

import "testing"

func TestRedisPingCheck(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				addr: "127.0.0.1:6379",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RedisPingCheck(tt.args.addr)
		})
	}
}
