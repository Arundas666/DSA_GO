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
	ll.head.next = &Node{data: 1}
	ll.head.next.next = &Node{data: 2}
	ll.head.next.next.next = &Node{data: 2}
	ll.head.next.next.next.next = &Node{data: 2}
	ll.head = RemoveDuplicates(ll.head)
	PrintFH(ll.head)

}
func RemoveDuplicates(head *Node) *Node {
	currentNode := head
	for currentNode != nil && currentNode.next != nil {
		if currentNode.data == currentNode.next.data {
			currentNode.next = currentNode.next.next
		} else {
			currentNode = currentNode.next
		}
	}
	return head
}

func PrintFH(head *Node) {
	if head == nil {
		fmt.Println("ntg to print")
	}
	curentNode := head
	for curentNode != nil {
		fmt.Print(curentNode.data)
		curentNode = curentNode.next
	}
}
