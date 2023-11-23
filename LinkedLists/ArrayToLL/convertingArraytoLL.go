package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func main() {
	var ll LinkedList
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	ll.head = ArraytoLinkedList(arr)
	PrintLinkedListFromHead(ll.head)
}
func ArraytoLinkedList(arr []int) *Node {
	var head, tail *Node

	for _, val := range arr {
		newNode := &Node{
			data: val,
			next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
	return head
}
func PrintLinkedListFromHead(head *Node) {
	fmt.Println("")
	for head != nil {
		fmt.Printf(" %d ", head.data)
		head = head.next
	}
	fmt.Println("")
}
