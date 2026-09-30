package main

import "testing"

func TestSquaresOfASortedArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "default case",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "same absolute nums",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "nothing",
			nums: []int{},
			want: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SquaresOfASortedArrays(test.nums)

			if !arrEqual(got, test.want) {
				t.Errorf(
					"SquaresOfASortedArrays(%v) = %v, want %v",
					test.nums,
					got,
					test.want,
				)
			}
		})
	}
}
