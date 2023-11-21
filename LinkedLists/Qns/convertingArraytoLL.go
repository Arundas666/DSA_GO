package Qns

import "fmt"

func (dll *DoubleLL) ArraytoLinkedList(arr []int) *Node {
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
