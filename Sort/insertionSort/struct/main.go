package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var arr = []person{
		{name: "zrun", age: 12}, {name: "s", age: 12},
	}
	a := sort(arr)
	fmt.Println(a)
}
func sort(nums []person) []person {
	for i := 0; i < len(nums); i++ {
		key := nums[i].name
		keyfull := nums[i]
		j := i - 1
		for j >= 0 && nums[j].name >key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = keyfull
	}
	return nums
}
