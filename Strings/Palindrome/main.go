package main

import "fmt"

func main() {
	Palindrome("aAa")
	PalindromeWithoutCapitalOsSmall("12321")

}
func Palindrome(str string) {

	count := 0
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			count = 1

		}

	}
	if count == 1 {
		fmt.Println("NOT A PALINDROME")
	} else {
		fmt.Println("PALINDROME")
	}

}

func PalindromeWithoutCapitalOsSmall(str string) {
	character := []rune(str)
	count := 0

	for i, j := 0, len(character)-1; i < len(character)/2; i, j = i+1, j-1 {
		if character[i] != character[j] && character[i] != character[j]-32 && character[i] != character[j]+32 {
			count = 1
		}
	}
	if count == 0 {
		fmt.Println("PALINDROME")
	} else {
		fmt.Println("NOT A PALINDROME")
	}
}
