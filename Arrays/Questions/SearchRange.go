package questions

import "fmt"

func SearchRange(nums []int, target int) []int {
	var start = -1
	var end = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && start == -1 {
			start = i
			fmt.Println(start, "start")
			for j := i; j < len(nums); j++ {
				if nums[j] != target {
					end = j - 1
					fmt.Println(end, "end")
				}
			}
		}
	}
	var a []int
	a = append(a, start, end)

	return a
}
