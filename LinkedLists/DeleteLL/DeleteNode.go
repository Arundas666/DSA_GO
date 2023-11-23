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
	ll.head = &Node{data: 1}
	ll.head.next = &Node{data: 2}
	ll.head.next.next = &Node{data: 4}
	ll.head.next.next.next = &Node{data: 4}
	ll.head.next.next.next.next = &Node{data: 5}
	ll.head = DeleteNode(ll.head, 1)
	PrintFromHdhead(ll.head)
}
func DeleteNode(head *Node, value int) *Node {
	if head != nil && head.data == value {
		return head.next
	}
	currentNode := head
	var prevNode *Node
	for currentNode.data != value {
		prevNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode != nil {
		prevNode.next = currentNode.next
	}
	return head
}

func PrintFromHdhead(head *Node) {
	if head == nil {
		fmt.Println("Nothing to print")
	}
	currentNode := head
	for currentNode != nil {
		fmt.Print(currentNode.data)
		currentNode = currentNode.next
	}
}
