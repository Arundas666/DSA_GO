package main

import "fmt"

func arraySum(arr []int, index int) int {
	if index < 0 {
		return 0
	}
	return arr[index] + arraySum(arr, index-1)
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	sum := arraySum(array, len(array)-1)
	fmt.Printf("Array: %v\nSum: %d\n", array, sum)
}
