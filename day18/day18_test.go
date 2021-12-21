package day18

import "testing"

func Test_magnitude(t *testing.T) {
	tests := []struct {
		name   string
		values []valueNode
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "Test Magnitude",
			values: parseLine("[[9,1],[1,9]]"),
			want:   129,
		},
		{
			name:   "Test Magnitude2",
			values: parseLine("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"),
			want:   3488,
		},
		{
			name:   "Test Magnitude3",
			values: parseLine("[[1,2],[[3,4],5]]"),
			want:   143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := magnitude(tt.values)
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
