package main

import "fmt"

func main() {
	var arr = []int{90, 2, 0, 100, 6, 4, 57, 1, 12, 43}
	fmt.Println("after sorting", quicksort(arr, 0, 9))

}
func quicksort(arr []int, strindx int, endidx int) []int {
	if strindx>=endidx{
		return nil
	}
	pivotIdx := strindx
	leftidx := strindx + 1
	rightidx := endidx

	for leftidx <= rightidx {
		if arr[leftidx] > arr[pivotIdx] && arr[rightidx] < arr[pivotIdx] {
			arr[leftidx], arr[rightidx] = arr[rightidx], arr[leftidx]
			leftidx++
			rightidx--
		}
		if arr[leftidx] <= arr[pivotIdx] {
			leftidx++
		}
		if arr[rightidx] >= arr[pivotIdx] {
			rightidx--
		}
	}
	arr[pivotIdx], arr[rightidx] = arr[rightidx], arr[pivotIdx]

	quicksort(arr, strindx, rightidx-1)
	quicksort(arr, rightidx+1, endidx)
	return arr

}
