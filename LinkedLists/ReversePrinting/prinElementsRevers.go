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
	ll.head.next.next = &Node{data: 3}
	ll.head.next.next.next = &Node{data: 4}
	PrintElementsInOrderRecursive(ll.head)
	fmt.Println("\nReverse order")
	PrintElementsInReverseOrderRecursive(ll.head)

}

func PrintElementsInOrderRecursive(head *Node) {
	if head != nil {
		fmt.Printf("%d ", head.data)
		PrintElementsInOrderRecursive(head.next)
	}
}
func PrintElementsInReverseOrderRecursive(head *Node) {
	if head != nil {
		PrintElementsInReverseOrderRecursive(head.next)
		fmt.Printf("%d ", head.data)
	}
}
