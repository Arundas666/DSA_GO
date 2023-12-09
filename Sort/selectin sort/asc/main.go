package main

import "fmt"

func main() {
	var arr = []int{4, 5, 3, 7, 2, 6}
	a := asc(arr)
	fmt.Println(a)

}
func asc(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}
