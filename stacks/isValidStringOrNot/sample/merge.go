package main

import "fmt"

func main() {
	var arr = []int{4, 6, 2, 1, 6, 1}
	fmt.Println(mergesort(arr))
}
func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	result := merge(left, right)
	return result
}
func merge(left []int, right []int) []int {
	i, j := 0, 0
	var result []int
	
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
