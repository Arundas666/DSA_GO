package main

import "fmt"

type queue struct {
	arr []int
}

func main() {
	q := queue{}
	q.enqeue(3)
	q.enqeue(4)
	q.enqeue(5)
	q.enqeue(6)
	fmt.Println(q.arr)
	q.dequeue()
	fmt.Println(q.arr)

}
func (q *queue) enqeue(element int) {
	q.arr = append(q.arr, element)
}
func (q *queue) dequeue() {
	if len(q.arr) == 0 {
		return
	}
	q.arr = q.arr[1:]

}
