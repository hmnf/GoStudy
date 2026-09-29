package main

import "testing"

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Unique in the start",
			s:    "leetcode",
			want: 0,
		},
		{
			name: "Unique in the middle",
			s:    "loveleetcode",
			want: 2,
		},
		{
			name: "No unique",
			s:    "aabb",
			want: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FirstUnique(test.s)

			if got != test.want {
				t.Errorf(
					"FirstUnique(%v) = %d; want %d",
					test.s,
					got,
					test.want,
				)
			}
		})
	}
}
