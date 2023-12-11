package main

import "fmt"

type queue struct {
	head *node
	end  *node
}
type node struct {
	data     int
	nextNode *node
}

func main() {
	var que queue
	que.enqueue(1)
	que.enqueue(2)
	que.enqueue(3)
	que.enqueue(4)
	print(&que)
	que.dequeue()
	print(&que)

}
func (q *queue) enqueue(element int) {
	newNode := &node{data: element,nextNode: nil}
	if q.head == nil {
		q.head = newNode
	} else {
		q.end.nextNode = newNode
	}
	q.end=newNode
}
func (q *queue) dequeue() {
	if q.head == nil {
		return
	}
	q.head = q.head.nextNode
}
func print(q *queue){
	if q.head==nil{
		fmt.Println("Nothing to print")
		return
	}
	currentNode:=q.head
	for currentNode!=nil{
		fmt.Print(currentNode.data)
		currentNode=currentNode.nextNode
	}
	fmt.Println(" ")
}
