package main

import "fmt"

type queue struct {
	ll linkedlist
}
type linkedlist struct {
	head *node
	size int
}
type node struct {
	data int

	nextNode *node
}

func (q *queue) enqueue(data int) {
	newNode := &node{data: data}
	if q.ll.head == nil {
		q.ll.head = newNode
		q.ll.size++
		return
	}
	currentNode := q.ll.head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}
	currentNode.nextNode = newNode
	q.ll.size++

}
func (q *queue) dequeue() int {
	if q.ll.head == nil {
		return 0
	}
	frstelement := q.ll.head.data
	q.ll.head = q.ll.head.nextNode
	return frstelement
}
func (q *queue) removeMid(mididx int, currentidx int) {
	if currentidx == mididx {
		mid := q.dequeue()
		fmt.Println("mid is", mid)
		return
	}
	temp := q.dequeue()
	q.removeMid(mididx, currentidx+1)
	q.enqueue(temp)
}
func print(q *queue) {
	if q.ll.head == nil {
		fmt.Println("Nothing to print")
		return
	}
	currentNode := q.ll.head
	for currentNode != nil {
		fmt.Print(currentNode.data)
		currentNode = currentNode.nextNode
	}
	fmt.Println(" ")
}
func main() {
	q := &queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(9)
	q.enqueue(8)
	q.enqueue(4)
	q.enqueue(5)
	mid := q.ll.size / 2
	q.removeMid(mid, 0)
	print(q)

}
