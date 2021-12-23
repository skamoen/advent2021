package day21

import (
	"testing"
)

func Test_d_Run(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "Test run",
			want:  412344,
			want1: 214924284932572,
		},
	}
	for _, tt := range tests {
		inputFile = "./input.txt"

		t.Run(tt.name, func(t *testing.T) {
			d := &d{}
			for i := 0; i < 1000; i++ {

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
