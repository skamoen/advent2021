package util

import "testing"

func TestUintBinaryToDecimal(t *testing.T) {
	type args struct {
		b []uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test binary conversion 1",
			args: args{b: []uint8{0, 0, 1}},
			want: 1,
		},
		{
			name: "Test binary conversion 2",
			args: args{b: []uint8{0, 1, 0}},
			want: 2,
		},
		{
			name: "Test binary conversion 3",
			args: args{b: []uint8{0, 1, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintBinaryToDecimal(tt.args.b); got != tt.want {
				t.Errorf("UintBinaryToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
