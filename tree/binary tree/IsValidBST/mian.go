package main

import "fmt"

type BinaryST struct {
	root *node
}
type node struct {
	data   int
	left   *node
	right  *node
	parent *node
}

func IsValidBST(root *node) bool {
	
	return checkingBST(root, nil, nil)
}
func checkingBST(node *node, min, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.data <= *min {
		return false

	}
	if max != nil && node.data >= *max {
		return false
	}
	val := node.data
	return checkingBST(node.left, min, &val) && checkingBST(node.right, &val, max)

}
func (b *BinaryST)addNode(data int){
	newNode:=node{data: data}
	if b.root==nil{
		b.root=&newNode
		return
	}
	currentNode:=b.root
	var parentNode *node
	for currentNode!=nil{
		parentNode=currentNode
		if currentNode.data>data{
			currentNode=currentNode.left
		}else{
			currentNode=currentNode.right
		}
	}
	if newNode.data>parentNode.data{
		parentNode.right=&newNode
	}else{
		parentNode.left=&newNode
	}
}
func main() {
	b := BinaryST{}
	b.addNode(6)
	b.addNode(3)
	b.addNode(7)
	b.addNode(8)
	
	
	f := IsValidBST(b.root)
	fmt.Println(f)
}
