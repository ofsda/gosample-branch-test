package sample

import "testing"

func TestSample(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want int
	}{
		{
			name: "test1",
			v:    0,
			want: 1,
		},
		{
			name: "test2",
			v:    1,
			want: 2,
		},
		{
			name: "test3",
			v:    2,
			want: 3,
		},
		{
			name: "test4",
			v:    3,
			want: 4,
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
