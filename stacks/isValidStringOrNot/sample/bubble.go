package main

import "fmt"

func main() {
	var arr = []int{4, 3, 2, 1, 5}
	bubblesort(arr)
	fmt.Println(arr)
}
func bubblesort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		swap := 0
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap = 1
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		if swap == 0 {
			break
		}
	}
}
