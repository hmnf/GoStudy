package main

import (
	"slices"
	"testing"
)

func TestMergeKSorted(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int
		want []int
	}{
		{
			name: "1",
			arr: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "2",
			arr: [][]int{
				{1, 10},
				{},
				{2, 3, 15},
				{4},
			},
			want: []int{1, 2, 3, 4, 10, 15},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mergeKSorted(test.arr)

			if !slices.Equal(got, test.want) {
				t.Errorf(
					"mergeKSorted(%v) = %v, want %v",
					test.arr,
					got,
					test.want,
				)
			}
		})
	}
}
