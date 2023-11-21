package Qns

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
