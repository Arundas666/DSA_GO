package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) insert(data int) {
	h.arr = append(h.arr, data)

}
func (h *maxHeap) insertHelper(position int) {
	for position > 0 {
		parentIdx := parent(position)
		if h.arr[position] > h.arr[parentIdx] {
			swap(h.arr, position, parentIdx)
			position = parentIdx
		} else {
			break
		}
	}
}
func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func display(arr []int) {
	for _, v := range arr {
		fmt.Print(v ," ")
	}
	fmt.Println(" ")
}
func (h *maxHeap) buildHeap(arr []int) {
	h.arr = arr
	for i := parent(len(h.arr) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}

}
func (h *maxHeap) shiftDown(currentIdx int) {

	endIdx := len(h.arr) - 1
	leftidx := leftChild(currentIdx)

	for leftidx <= endIdx {
		var idxToswap int
		rightidx := rightChild(currentIdx)
		if rightidx <= endIdx && h.arr[rightidx] > h.arr[leftidx] {
			idxToswap = rightidx
		} else {
			idxToswap = leftidx
		}
		if h.arr[idxToswap] > h.arr[currentIdx] {
			swap(h.arr, idxToswap, currentIdx)
			currentIdx = idxToswap
            leftidx = leftChild(currentIdx)

		} else {
			return
		}
	}
}
func leftChild(i int) int {
	return 2*i + 1
}
func rightChild(i int) int {
	return 2*i + 2
}
func parent(i int) int {
    return (i - 1) / 2
}
func (h *maxHeap)delete()int{
	if h.arr==nil{
		return -1
	}
	maxValue:=h.arr[0]
	h.arr[0]=h.arr[len(h.arr)-1]
	h.arr=h.arr[:len(h.arr)-1]
	h.shiftDown(0)

	return maxValue
}
func (h *maxHeap)heapSort()[]int{
	for i:=len(h.arr)-1;i>=0;i--{
		h.arr[0],h.arr[i]=h.arr[i],h.arr[0]
		h.shiftDown(0)
	}
	return h.arr
}
func main() {
	var arr = []int{11, 2, 7, 4,12, 5, 9,10}
    h := &maxHeap{}
    h.buildHeap(arr)

    fmt.Println("Original array:")
    display(h.arr)

    h.heapSort()

    fmt.Println("Sorted array:")
    display(h.arr)


}