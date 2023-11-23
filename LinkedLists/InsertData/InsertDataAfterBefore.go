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
	ll.head = InsertNodeAfterValue(ll.head, 1, 5)
	PrintFH(ll.head)
	println("")
	ll.head = InsertNodeBeforeValue(ll.head, 1, 4)
	PrintFH(ll.head)

}
func InsertNodeAfterValue(head *Node, value int, newData int) *Node {
	newNode := &Node{
		data: newData,
	}
	if head == nil {
		return newNode
	}
	currentNode := head
	for currentNode.data != value {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode
	return head
}
func InsertNodeBeforeValue(head *Node, value int, newData int) *Node {
	newNode := &Node{
		data: newData,
	}
	if head == nil {
		return newNode
	}
	if head.data==value{
		newNode.next=head
		head=newNode
	}
	currentNode := head
	var prevNod *Node
	for currentNode.data != value {
		prevNod = currentNode
		currentNode = currentNode.next
	}
	prevNod.next = newNode
	newNode.next = currentNode
	return head
}

func PrintFH(head *Node) {
	if head == nil {
		fmt.Println("Ntg ot print")
	}
	curentNode := head
	for curentNode != nil {
		fmt.Print(curentNode.data)
		curentNode = curentNode.next
	}
}
