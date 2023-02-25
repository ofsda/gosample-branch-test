package sample

import "testing"

func TestSample(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want int
	}{
		{
			name: "test2",
			v:    1,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sample(tt.v); got != tt.want {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
