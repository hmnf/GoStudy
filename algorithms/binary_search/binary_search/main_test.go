package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "1",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "3",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "4",
			nums:   []int{5},
			target: -5,
			want:   -1,
		},
		{
			name:   "5",
			nums:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:   "6",
			nums:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Search(test.nums, test.target)

			if got != test.want {
				t.Errorf(
					"Search %v = %d, want %d",
					test.name,
					got,
					test.want,
				)
			}
		})
	}
}
