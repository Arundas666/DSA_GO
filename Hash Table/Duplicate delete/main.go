package main

import "fmt"

func main() {
	var arr = []int{2, 2, 3, 4, 5, 6, 5, 3, 2, 1}
	fmt.Println(duplicateDelete(arr))
}
func duplicateDelete(arr []int) []int {
	newTable := make(map[int]int)
	for i, nums := range arr {
		if _, ok := newTable[nums]; ok {
			//delete from array
			for j := i; j < len(arr)-1; j++ {
				arr[j] = arr[j+1]
			}
			arr = arr[:len(arr)-1]
		}
		newTable[nums] = i
	}
	return arr
}
