package main

import "fmt"

func main() {
	// var arr = []int{89, 4, 6, 4, 3, 6, 2, 1, 0}
	// fmt.Println(mergeSort(arr))
	var arr = []int{2, 4, 1, 6, 3, 8, 2}
	fmt.Println(mergeSort(arr))

}
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)

}
func merge(left []int, right []int) []int {

	var result = make([]int, 0, len(left)+len(right))
	i:=0
	j:=0
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
