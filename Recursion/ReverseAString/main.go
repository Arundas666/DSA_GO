package main

import "fmt"

func main() {
	result := Reverse("ARUnmnjknN")
	fmt.Println("Resuklt is ", result)
}
func Reverse(str string) string {
	if str == "" {
		return ""
	}
	return Reverse(str[1:]) + string(str[0])

}
