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

func TestUnique(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{-1, -2, 0, 1, 10, 22, 22, 22, 10, -2}, []int{-1, -2, 0, 1, 10, 22}},
	}

	for _, test := range tests {
		got := unique(test.nums)

		if !arrEqual(got, test.want) {
			t.Errorf(
				"unique(%v) = %v; want %v",
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
