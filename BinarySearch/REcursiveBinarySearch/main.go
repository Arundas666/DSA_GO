package main

import "fmt"

// Implement the binary search algorithm using a recursive approach.
// Compare its performance with the iterative version.
func main() {
	array := []int{2, 5, 8, 12, 165, 23}
	targetElement := 12
	result := Recursive(array, 0, 5, targetElement)
	if result == -1 {
		fmt.Println("NOT FOUND")
	} else {
		fmt.Println("found in ", result)
	}
}
func Recursive(arr []int, low int, high int, target int) int {
	if low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			return Recursive(arr, mid+1, high, target)
		}
		if arr[mid] > target {
			return Recursive(arr, low, mid-1, target)
		}
	}
	return -1
}
