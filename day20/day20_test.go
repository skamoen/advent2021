package day20

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
			want:  5203,
			want1: 18806,
		},
	}
	for _, tt := range tests {
		inputFile = "./input.txt"

		t.Run(tt.name, func(t *testing.T) {
			d := &d{}
			for i := 0; i < 10; i++ {

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
