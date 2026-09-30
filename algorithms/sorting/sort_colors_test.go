package main

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "3 nums",
			nums: []int{1, 2, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "1's on two sides",
			nums: []int{1, 2, 2, 2, 2, 0, 0, 0, 1, 1},
			want: []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SortColors(test.nums)

			if !slices.Equal(test.nums, test.want) {
				t.Errorf(
					"SortColors() = %v, want %v",
					test.nums,
					test.want,
				)
			}
		})
	}
}
