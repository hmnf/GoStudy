package main

import "testing"

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "default case",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "nums 2 < nums 1",
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "nothing in nums 2",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "nothing in nums 1",
			nums1: []int{0},
			m:     1,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MergeSortedArrays(test.nums1, test.m, test.nums2, test.n)

			if !arrEqual(got, test.want) {
				t.Errorf(
					"MergeSortedArrays(%v, %v, %v, %v) = %v, want %v",
					test.nums1,
					test.m,
					test.nums2,
					test.n,
					got,
					test.want,
				)
			}
		})
	}
}
