package main

import (
	questions "arrays/Questions"
	"fmt"
)

func main() {

	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	isValid := questions.IsValidSudoku(sudoku)
	if isValid {
		fmt.Println("It is valid")
	} else {
		fmt.Println("Its not valid")
	}
	var nums = []int{5, 7, 7, 8, 8, 10}
	var nums2 = []int{5, 7, 7, 9, 9, 10}
	r := questions.SearchRange(nums, 8)
	fmt.Println(r)

	uniqueCount := questions.RemoveDuplicates(nums)
	fmt.Println(nums, uniqueCount)
	intArray := questions.Intersection(nums, nums2)
	fmt.Println("Intersection array is", intArray)

}
