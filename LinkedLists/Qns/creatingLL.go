package Qns

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
func (ll *LinkedList) AppendLL(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) DisplayLL() {
	current := ll.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

}
