package main

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic",
			nums: []int{5, 2, 4, 6, 1, 3},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "already sorted",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "reverse sorted",
			nums: []int{4, 3, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "same nums",
			nums: []int{5, 5, 5},
			want: []int{5, 5, 5},
		},
		{
			name: "nums < 0",
			nums: []int{-2, -8, -1, 2, 3, 6, 0},
			want: []int{-8, -1, -1, 0, 2, 3, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			insertionSort(test.nums)

			if !arrEqual(test.nums, test.want) {
				t.Errorf(
					"insertionSort() = %v, want %v",
					test.nums,
					test.want,
				)
			}
		})
	}
}
