package main

import "fmt"

func main() {
	v, c := CountVowelsAndConsonants("ARUN. .a a aaaa DA S")
	fmt.Println("Vowels:", v)
	fmt.Println("Consonant:", c)
}
func CountVowelsAndConsonants(str string) (int, int) {
	var vowels, consonant int
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u','A','E','I','O','U':
			vowels++
		case ' ', '.', '!', '$':
		default:
			consonant++

		}
	}

	return vowels, consonant
}
