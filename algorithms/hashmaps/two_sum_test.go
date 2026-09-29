package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "answer in middle",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "duplicate numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   []int{},
		},
		{
			name:   "empty slice",
			nums:   []int{},
			target: 5,
			want:   []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.nums, test.target)

			if !arrEqual(got, test.want) {
				t.Errorf(
					"TwoSum(%v, %d) = %v; want %v",
					test.nums,
					test.target,
					got,
					test.want,
				)
			}
		})
	}
}

func arrEqual(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}
