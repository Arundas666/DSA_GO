package main

import "fmt"

func main() {
	var arr = []int{4, 5, 3, 7, 2, 6}
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	for i:= range arr {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				arr[j], arr[minIdx] = arr[minIdx], arr[j]
			}

		}
	}

}
