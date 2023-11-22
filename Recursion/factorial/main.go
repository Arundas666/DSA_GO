package main

import "fmt"

func main() {
	result := factorial(4)
	fmt.Println("Factorial is ", result)

}

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}
