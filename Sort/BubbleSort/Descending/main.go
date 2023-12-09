package main

import "fmt"

func main() {
	var arr = []int{1, -1, -2,-3, 0}
	desc(arr)
	fmt.Println(arr)


}
func desc(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		swap := 0
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] < nums[j+1] {
				swap = 1
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if swap == 0 {
			break
		}
	}
	return nums
}
