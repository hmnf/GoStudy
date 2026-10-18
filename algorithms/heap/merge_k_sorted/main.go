package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
	x     int
	y     int
}

type ArrHeap []Item

func (h ArrHeap) Len() int {
	return len(h)
}

func (h ArrHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h ArrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ArrHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *ArrHeap) Pop() any {
	old := *h
	n := len(*h)

	item := old[n-1]

	old[n-1] = Item{}
	*h = old[:n-1]

	return item
}

func main() {
	fmt.Println(mergeKSorted([][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}))
}

func mergeKSorted(arrays [][]int) []int {
	var arrHeap ArrHeap
	heap.Init(&arrHeap)
	var res []int

	for i := 0; i < len(arrays); i++ {
	}

	for arrHeap.Len() > 0 {
		res = append(res, heap.Pop(&arrHeap).(Item).value)
	}
	return res
}
