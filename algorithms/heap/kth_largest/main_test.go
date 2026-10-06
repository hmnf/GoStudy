package main

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := findKthLargest(test.nums, test.k)

			if got != test.want {
				t.Errorf(
					"findKthLargest(%v, %d) = %d, want %d",
					test.nums,
					test.k,
					got,
					test.want,
				)
			}
		})
	}
}
