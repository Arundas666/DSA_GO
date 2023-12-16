package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index] > h.array[parentIndex] {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	max := h.array[0]
	lastIndex := len(h.array) - 1

	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifyDown(0)

	return max, nil
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1

	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largest := index

		if leftChildIndex <= lastIndex && h.array[leftChildIndex] > h.array[largest] {
			largest = leftChildIndex
		}

		if rightChildIndex <= lastIndex && h.array[rightChildIndex] > h.array[largest] {
			largest = rightChildIndex
		}

		if largest != index {
			h.swap(index, largest)
			index = largest
		} else {
			break
		}
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	maxHeap := NewMaxHeap()

	values := []int{10, 5, 12, 9, 5, 7, 3, 1}

	for _, value := range values {
		maxHeap.Insert(value)
	}

	for _,v:=range maxHeap.array{
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}

