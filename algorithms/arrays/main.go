package main

import "fmt"

func main() {
	nums1 := make([]int, 6)
	nums2 := make([]int, 3)
	nums1[0] = 1
	nums1[1] = 2
	nums1[2] = 3
	nums2[0] = 2
	nums2[1] = 5
	nums2[2] = 6
	fmt.Println(
		MergeSortedArrays(nums1, 3, nums2, 3),
	)
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func max(nums []int) int {
	mx := nums[0]
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}

	return mx
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

func onlyEven(nums []int) []int {
	var rs []int
	for _, num := range nums {
		if num%2 == 0 {
			rs = append(rs, num)
		}
	}
	return rs
}

func reverse(nums []int) []int {
	l := len(nums)
	rs := make([]int, l)

	for i, num := range nums {
		rs[l-i-1] = num
	}

	return rs
}

/*////////// Arrays ////////////

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// ////////// Slices ///////////////////
func main2() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

///////////////// Slices 2 ////////////////

func main3() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

////////// Tic Tac Toe //////////

func main4() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// ///////////////////  Maps //////////////////////////
type Vertex struct {
	Lat, Long float64
}

func main5() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/////////////////// Maps 2 ///////////////////////////

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main6() {
	fmt.Println(m)
}

// ///////////////// Word Count /////////////////////////////
var res map[string]int

func WordCount(s string) map[string]int {
	res = make(map[string]int)
	for _, v := range strings.Fields(s) {
		res[v]++
	}

	return res
}

func main7() {
	// wc.Test(WordCount)
}
*/
