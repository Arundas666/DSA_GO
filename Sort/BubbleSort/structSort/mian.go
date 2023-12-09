package main

import "fmt"

type Presons struct {
	Name string
	Age  int
}

func main() {
	a := []Presons{
		{Name: "avv",
			Age: 12},
		{
			Name: "b",
			Age:  1,
		},
		{Name: "avvvvv", Age: 3},
	}
	fmt.Println(a)
	sortAplabaticaly(a)
	fmt.Println(a)
}
func sortAplabaticaly(p []Presons) {
	for i := 0; i < len(p)-1; i++ {
		swap:=0
		for j := 0; j < len(p)-i-1; j++ {
			if p[j].Name > p[j+1].Name {
				swap=1
				p[j], p[j+1] = p[j+1], p[j]
			}
		}
		if swap==0{
			break
		}

	}

}
