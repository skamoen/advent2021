package day16

import (
	"testing"
)

func Test_d_Run(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "Test run",
			want:  940,
			want1: 13476220616073,
		},
	}
	inputFile = "./input.txt"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &d{}
			for i := 0; i < 10000; i++ {
				got, got1 := d.Run()
				if got != tt.want {
					t.Errorf("Run() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
				}
			}
		})
	}
}
