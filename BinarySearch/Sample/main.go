package main

import "fmt"

func main() {
	array := []int{2, 5, 8, 12, 165, 23}
	targetElement := 16
	result := BinarySearch(array, targetElement)
	if result == -1 {
		fmt.Println("Elemetn not found")
	} else {
		fmt.Println("Element found in index ", result)
	}
}
func BinarySearch(arr []int, target int) int {
	l := 0
	e := len(arr) - 1

	for l <= e {
		mid := (l + e) / 2

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1

		}
		if arr[mid] > target {
			e = mid - 1
		}
	}
	return -1

}
