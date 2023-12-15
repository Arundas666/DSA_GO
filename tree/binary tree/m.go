package main

import "fmt"

func main() {
	// var arr = []int{1, 1, 1, 2}
var arr = []int{0,0,1,1,1,2,2,3,3,4}
	k := removeDuplicates(arr)
	fmt.Println(arr[:k])

}
func removeDuplicates(nums []int) int {
	k := 1
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				y := i + 1
				for x := j; x < len(nums); x++ {
					nums[y] = nums[x]
					y++
				}
				k++
				n--

				break
			}
		}
	}
	return k
}
