package main

import "fmt"

type Node struct {
	data   int
	lChild *Node
	rChild *Node
	parent *Node
}
type BinarySearchTree struct {
	root *Node
}

func main() {

	b := &BinarySearchTree{}
	b.root=addNode(b.root,5)
	b.root=addNode(b.root,6)
	b.root=addNode(b.root,6)
	b.root=addNode(b.root,7)
	b.root=addNode(b.root,8)
	
	b.printInorder(b.root)
	b.contain(9)

}
func  addNode(root *Node,data int) (*Node){
	if root == nil {
		root = &Node{data: data,lChild: nil,rChild: nil}	
		return root
	}
	
	
		if data > root.data {
			root.rChild=addNode(root.rChild,data)
		} else {
			root.rChild=addNode(root.lChild,data)
		}
		return  root
}
func (b *BinarySearchTree) contain(data int) {
	temp := b.root
	if temp == nil {
		fmt.Println("Tree is empty")
		return
	}
	for temp != nil {

		if data < temp.data {
			temp = temp.lChild

		} else if data > temp.data {
			temp = temp.rChild
		} else {

			fmt.Println("MAtch Found")
			return
		}

	}
	fmt.Println("MAtch not found")
	return

}
func (b *BinarySearchTree)printInorder(node *Node){
	if node==nil{
		
		return
	}
	b.printInorder(node.lChild)
	fmt.Print(node.data)
	b.printInorder(node.rChild)

}
func (b *BinarySearchTree)PreOrder(node *Node){
	if node==nil{
		return
	}
	fmt.Print(node.data)
	b.PreOrder(node.lChild)
	b.PreOrder(node.rChild)
}