package main

import (
	"LL/Qns"
	"fmt"
)

func main() {
	newLL := Qns.NewLinkedList()
	newLL.AppendLL(12)
	newLL.AppendLL(13)
	newLL.AppendLL(122)
	newLL.AppendLL(1)
	newLL.DisplayLL()

	// Dll := Qns.NewDoubleLL()
	// Dll.AddingIntoDLL(1)
	// Dll.AddingIntoDLL(2)
	// Dll.AddingIntoDLL(3)
	// Dll.AddingIntoDLL(4)
	// Dll.AddingIntoDLL(5)
	// Dll.AddingIntoDLL(6)
	// Dll.DisplayDataForward()
	// Dll.DisplayDataReverse()
	// arr := []int{10, 20, 30, 40, 50}
	// headd := Dll.ArraytoLinkedList(arr)
	// Qns.PrintLinkedListFromHead(headd)
	var head *Qns.Node
	// head = Qns.AddNodeAtEnd(head, 102)
	// head = Qns.AddNodeAtEnd(head, 202)
	// head = Qns.AddNodeAtEnd(head, 402)
	// head = Qns.AddNodeAtEnd(head, 302)
	// Qns.PrintLinkedListFromHead(head)
	head = Qns.AddNodeAtBeginnning(head, 900)
	head = Qns.AddNodeAtBeginnning(head, 899)
	head = Qns.AddNodeAtBeginnning(head, 898)
	head = Qns.AddNodeAtEnd(head, 9008)

	// Qns.PrintLinkedListFromHead(head)
	// head = Qns.DeleteNode(head, 900)
	// Qns.PrintLinkedListFromHead(head)
	Qns.PrintLinkedListFromHead(head)
	head = Qns.InsertNodeAfterValue(head, 899, 777)
	head = Qns.InsertNodeBeforeValue(head, 900, 777)
	Qns.PrintLinkedListFromHead(head)
	fmt.Println("PRinting in order is ")
	Qns.PrintElementsInOrderRecursive(head)
	fmt.Println("PRinting in Revesrse order is ")

	Qns.PrintElementsInReverseOrderRecursive(head)
	head = Qns.RemoveDuplicates(head)
	fmt.Println("After deleting duplicates")
	Qns.PrintLinkedListFromHead(head)

}
