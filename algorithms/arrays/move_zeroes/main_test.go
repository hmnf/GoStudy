package main

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "only one zero",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "one num and zeroes",
			nums: []int{0, 0, 0, 1},
			want: []int{1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moveZeroes(test.nums)

			if !slices.Equal(test.nums, test.want) {
				t.Errorf(
					"MoveZeroes() = %v, want = %v",
					test.nums,
					test.want,
				)
			}
		})
	}
}
