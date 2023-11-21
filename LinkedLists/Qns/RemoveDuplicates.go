package Qns

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
