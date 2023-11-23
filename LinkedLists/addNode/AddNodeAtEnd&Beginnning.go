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
	ll.head = AddNodeAtEnd(ll.head, 1)
	ll.head = AddNodeAtEnd(ll.head, 2)
	ll.head = AddNodeAtEnd(ll.head, 3)
	ll.head = AddNodeAtEnd(ll.head, 4)
	ll.head = AddNodeAtEnd(ll.head, 5)
	PrintFromHead(ll.head)
	ll.head = AddNodeAtBeginnning(ll.head, 9)
	fmt.Println(" ")
	PrintFromHead(ll.head)
}

func AddNodeAtEnd(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
		next: nil,
	}
	if head == nil {
		return newNode
	}
	curentNode := head
	for curentNode.next != nil {
		curentNode = curentNode.next
	}
	curentNode.next = newNode
	return head
}
func AddNodeAtBeginnning(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
		next: head,
	}
	return newNode

}

func PrintFromHead(head *Node) {
	if head == nil {
		fmt.Println("Nothing to print")
	}
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%d", currentNode.data)
		currentNode = currentNode.next
	}
}
