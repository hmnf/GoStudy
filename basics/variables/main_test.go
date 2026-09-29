package main

import "testing"

func TestSumDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{1234, 10},
		{105, 6},
		{5, 5},
		{0, 0},
	}

	for _, test := range tests {
		got := sumDigits(test.num)

		if got != test.want {
			t.Errorf(
				"sumDigits(%d) = %d; want %d",
				test.num,
				got,
				test.want,
			)
		}
	}
}

func TestMaxOfThree(t *testing.T) {
	tests := []struct {
		nums []float64
		want float64
	}{
		{[]float64{5, 12, 8}, 12},
		{[]float64{100, 5, 3}, 100},
		{[]float64{1, 2, 100}, 100},
		{[]float64{-5, -2, -10}, -2},
	}

	for _, test := range tests {
		got := maxOfThree(test.nums[1], test.nums[2], test.nums[3])

		if got != test.want {
			t.Errorf(
				"maxOfThree(%d) = %d, want %d",
				test.nums,
				got,
				test.want,
			)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
}

func TestCalc(t *testing.T) {
}
