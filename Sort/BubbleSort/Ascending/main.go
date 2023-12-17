package main

import "fmt"

func main() {

	var arr = []int{66, 43, 44, 45, 4}
	asc(arr)
	fmt.Println(arr)

}

func asc(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		swap := 0
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j+1] < nums[j] {
				swap = 1
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if swap == 0 {
			break
		}
	}
	return nums
}
