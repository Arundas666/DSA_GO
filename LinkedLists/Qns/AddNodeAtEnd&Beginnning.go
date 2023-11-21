package Qns

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
