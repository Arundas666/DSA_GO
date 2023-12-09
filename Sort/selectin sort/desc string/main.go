package main

import "fmt"

func main() {
	var arr = []string{"aaa", "a", "asfa", "z", "adsads"}
	a := desc(arr)
	fmt.Println(a)
}
func desc(str []string) []string {
	for i := 0; i < len(str)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(str); j++ {
			if len(str[j]) > len(str[minIndex]) {
				minIndex = j
			}
		}
		str[i], str[minIndex] = str[minIndex], str[i]
	}
	return str
}
