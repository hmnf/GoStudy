package main

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "true",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "> 2 elems",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "false",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			k:    3,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := containsNearbyDuplicate(test.nums, test.k)

			if got != test.want {
				t.Errorf(
					"containsNearbyDuplicate(%v, %d) = %v, want %v",
					test.nums,
					test.k,
					got,
					test.want,
				)
			}
		})
	}
}
