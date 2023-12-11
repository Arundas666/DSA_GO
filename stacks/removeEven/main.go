package main

import "fmt"

type stack struct {
	list []int
}

func main() {
	s := &stack{}
	s.push(1)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(2)
	fmt.Println(s.list)
	s.removeEvenNumbers()
	fmt.Println(s.list)
	s.pop()
	fmt.Println(s.list)

}
func (s *stack) push(a int) {
	s.list = append(s.list, a)
}
func (s *stack) pop() int {
	if len(s.list) == 0 {
		fmt.Println("array is empty")
		return 0
	}
	lastelement := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return lastelement
}
func (s *stack) removeEvenNumbers() {
	if len(s.list) == 0 {
		return
	}
	element := s.pop()
	s.removeEvenNumbers()
	if element%2 != 0 {
		s.push(element)
	}
}
