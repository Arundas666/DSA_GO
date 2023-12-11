package main

import "fmt"

func main() {

	fmt.Println(isValid("[][()]"))

}
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		if char == '[' || char == '(' || char == '{' {
			stack = append(stack, char)
		} else if char == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else if char == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if char == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
	
}
