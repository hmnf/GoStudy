package main

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "2",
			nums: []int{4, 4, 4, 6, 6, 7, 7, 7, 7, 8},
			k:    2,
			want: []int{7, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := topKFrequent(test.nums, test.k)

			if !slices.Equal(got, test.want) {
				t.Errorf(
					"topKFrequent(%v, %d) = %v, want %v",
					test.nums,
					test.k,
					got,
					test.want,
				)
			}
		})
	}
}
