package Qns

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
