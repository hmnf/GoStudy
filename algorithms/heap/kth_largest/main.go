package main

import (
	"container/heap"
	"fmt"
)

type NumsHeap []int

func (h NumsHeap) Len() int {
	return len(h)
}

func (h NumsHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h NumsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NumsHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *NumsHeap) Pop() any {
	old := *h
	n := len(old)

	num := old[n-1]

	*h = old[:n-1]

	return num
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	var nHeap NumsHeap
	heap.Init(&nHeap)

	for _, num := range nums {
		heap.Push(&nHeap, num)
	}

	for k > 1 {
		heap.Pop(&nHeap)
	}

	return heap.Pop(&nHeap).(int)
}
