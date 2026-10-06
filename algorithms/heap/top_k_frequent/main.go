package main

import "container/heap"

type Item struct {
	num   int
	count int
}

type FreqHeap []Item

func (h FreqHeap) Len() int {
	return len(h)
}

func (h FreqHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h FreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreqHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *FreqHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]

	old[n-1] = Item{}
	*h = old[:n-1]

	return item
}

func main() {
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	var freqHeap FreqHeap
	heap.Init(&freqHeap)
	for num, count := range freq {
		heap.Push(
			&freqHeap,
			Item{
				num:   num,
				count: count,
			},
		)
	}
	var res []int
	for k > 0 {
		res = append(res, heap.Pop(&freqHeap).(Item).num)
		k--
	}

	return res
}
