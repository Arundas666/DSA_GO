package main

import "fmt"

func main() {
	var arr = []int{2, 5, 3, 5, 3, 8, 7, 5, 7, 1, 3}
	fmt.Println(twosumgivingAllSums(arr, 8))
	fmt.Println(twosum(arr, 8))

}
func twosum(arr []int, target int) []int {
	newtable := make(map[int]int)
	for i, nums := range arr {
		compliment := target - nums
		if j, ok := newtable[compliment]; ok {
			return []int{i, j}
		}
		newtable[nums] = i
	}
	return nil
}
func twosumgivingAllSums(arr []int, target int) [][]int {
	var result [][]int
	var newtable = make(map[int]int)
	for i, nums := range arr {
		compliment := target - nums
		if j, found := newtable[compliment]; found {
			result = append(result, []int{arr[i], arr[j]})
		}
		newtable[nums] = i
	}
	return result

}
