package main

import "fmt"

func multipleElementSearch(arr []int) []int {
	occurrences := make(map[int]int)
	var result []int

	// Count occurrences of each element
	for _, num := range arr {
		occurrences[num]++
	}

	// Add elements with more than one occurrence to the result slice
	for num, count := range occurrences {
		if count > 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	array := []int{16, 5, 5, 16, 16, 23, 16}

	resultArray := multipleElementSearch(array)
	if len(resultArray) > 0 {
		fmt.Println("Multiple occured elements are ", resultArray)
	} else {
		fmt.Println("Array is having unique elements only")
	}

}
