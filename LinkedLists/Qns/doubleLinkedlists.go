package Qns

import "fmt"

type NodeDLL struct {
	data int
	next *NodeDLL
	prev *NodeDLL
}
type DoubleLL struct {
	head *NodeDLL
	tail *NodeDLL
}

func NewDoubleLL() *DoubleLL {
	return &DoubleLL{}
}

func (dll *DoubleLL) AddingIntoDLL(data int) {
	newNode := NodeDLL{data: data, next: nil, prev: nil}
	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
		return
	} else {
		newNode.prev = dll.tail
		dll.tail.next = &newNode
		dll.tail = &newNode
	}
}
func (dll *DoubleLL) DisplayDataForward() {
	for dll.head != nil {
		fmt.Printf("%d ", dll.head.data)
		dll.head = dll.head.next
	}
	fmt.Println("")
}
func (dll *DoubleLL) DisplayDataReverse() {
	for dll.tail != nil {
		fmt.Printf("%d ", dll.tail.data)
		dll.tail = dll.tail.prev
	}
}
