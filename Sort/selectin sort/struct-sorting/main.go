package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var arr = []person{{name: "hey", age: 9}, {name: "aa", age: 2}, {name: "arrr", age: 9}}

	a := asc(arr)
	fmt.Println(a)
}
func asc(p []person) []person {
	for i := 0; i < len(p)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(p); j++ {
			if p[j].name < p[minIndex].name {
				minIndex = j
			}
		}
		p[i], p[minIndex] = p[minIndex], p[i]
	}
	return p
}
