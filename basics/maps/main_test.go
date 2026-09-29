package main

import "testing"

func TestFrequency(t *testing.T) {
	tests := []struct {
		nums []int
		want map[int]int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, map[int]int{1: 1, 2: 2, 3: 3}},
		{[]int{}, map[int]int{}},
		{[]int{-1, -1, 0, -1}, map[int]int{-1: 3, 0: 1}},
	}

	for _, test := range tests {
		got := frequency(test.nums)

		if !mapsEqual(got, test.want) {
			t.Errorf(
				"frequency(%v) = %v; want %v",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func mapsEqual(got, want map[int]int) bool {
	if len(got) != len(want) {
		return false
	}

	for key, value := range got {
		wantValue, ok := want[key]
		if !ok || value != wantValue {
			return false
		}
	}

	return true
}
