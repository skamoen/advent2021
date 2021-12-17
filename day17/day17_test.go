package day17

import (
	"testing"
	"time"
)

func Test_d_Run(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "Test Shooting",
			want:  2701,
			want1: 1070,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			d := &d{}
			for time.Since(start) < time.Second*5 {
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
