package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h myHeap) printHeap() {
	n := 1
	levelCount := 1
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}

func main() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap)

	for _, value := range data {
		aheap.Push(value)
	}

	aheap.printHeap()

	heap.Init(aheap)

	fmt.Println("aheap:", *aheap)

	aheap.printHeap()
}
