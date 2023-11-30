package main

import "fmt"

func ReverseString(str string) string {
	characters := []rune(str)
	for i, j := 0, len(characters)-1; i < len(characters)/2; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
	str = string(characters)
	return str
}
func main() {
	reversed := ReverseString("HELLO")
	fmt.Println(reversed)
}
