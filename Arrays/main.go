package main

import (
	questions "array/intersection"
	remv "array/removeDuplicate"
	sear "array/searchrange"
	"fmt"
)

func main() {

	var nums = []int{5, 7, 7, 8, 8, 10}
	var nums2 = []int{5, 7, 7, 9, 9, 10}
	r := sear.SearchRange(nums, 8)

	fmt.Println("search range ", r)

	uniqueCount := remv.RemoveDuplicates(nums)
	fmt.Println("unique coubnt is ", uniqueCount)
	intArray := questions.Intersection(nums, nums2)
	fmt.Println("Intersection array is", intArray)

}
