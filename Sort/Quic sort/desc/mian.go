package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 5, 4, 6, 7, 8, 9}
	fmt.Println("desc order is", quick(arr, 0, len(arr)-1))
}
func quick(arr []int, strt int, end int) []int {
	if strt >= end {
		return nil

	}
	pivot := strt
	left := strt+1
	right := end 
	for left <= right {

		if arr[left] < arr[pivot] && arr[right] > arr[pivot] {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
		if arr[left] >= arr[pivot] {
			left++
		}
		if arr[right] <= arr[pivot] {
			right--
		}
	}
	arr[pivot], arr[right] = arr[right], arr[pivot]
	quick(arr, strt, right-1)
	quick(arr, right+1, end)
	return arr

}
