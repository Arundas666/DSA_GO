package Qns

import "fmt"

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
