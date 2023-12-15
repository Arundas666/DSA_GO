package main

import (
	"fmt"
	"math"
)

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
	b.root = addNode(b.root, 10)
	b.root = addNode(b.root, 5)
	b.root = addNode(b.root, 6)
	b.root = addNode(b.root, 7)
	b.root = addNode(b.root, 4)
	b.root = addNode(b.root, 3)
	b.root = addNode(b.root, 8)
	b.root = delete(b.root, 8)

	b.printInorder(b.root)
	fmt.Println(" ")
	b.PreOrder(b.root)
	fmt.Println(" ")
	b.postOrder(b.root)
	fmt.Println(" ")
	b.contain(7)
	fmt.Println("CLOSEST OF 1 is ", b.closest(1))

	fmt.Println("isBST", isBST(b.root, 1, 10))
	b.printInorderWithConstraints(b.root,5,7)
}
func addNode(root *Node, data int) *Node {
	if root == nil {
		root = &Node{data: data, lChild: nil, rChild: nil}
		return root
	}

	if data > root.data {
		root.rChild = addNode(root.rChild, data)
	} else {
		root.lChild = addNode(root.lChild, data)
	}
	return root
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
func delete(root *Node, data int) *Node {

	if root == nil {
		return root
	}
	if root.data < data {
		root.rChild = delete(root.rChild, data)
	} else if root.data > data {
		root.lChild = delete(root.lChild, data)

	} else {
		if root.lChild == nil {
			return root.rChild
		} else if root.rChild == nil {
			return root.lChild
		} else {
			c := getSuccesor(root)
			root.data = c.data
			root.rChild = delete(root.rChild, c.data)
		}

	}
	return root

}
func getSuccesor(c *Node) *Node {
	c = c.rChild
	for c != nil && c.lChild != nil {
		c = c.lChild

	}
	return c

}
func (b *BinarySearchTree) printInorder(node *Node) {
	if node == nil {

		return
	}
	b.printInorder(node.lChild)
	fmt.Print(node.data)
	b.printInorder(node.rChild)

}
func (b *BinarySearchTree) PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data)
	b.PreOrder(node.lChild)
	b.PreOrder(node.rChild)
}
func (b *BinarySearchTree) postOrder(node *Node) {
	if node == nil {
		return
	}
	b.postOrder(node.lChild)
	b.postOrder(node.rChild)
	fmt.Print(node.data)
}
func (b *BinarySearchTree) closest(data int) int {
	if b.root == nil {
		return -1
	}
	minDiff := math.MaxInt

	var currentElement int
	for b.root != nil {
		currentDiff := math.Abs(float64(b.root.data) - float64(data))
		if currentDiff < float64(minDiff) {
			minDiff = int(currentDiff)
			currentElement = b.root.data
		}
		if data < b.root.data && b.root.lChild != nil {
			b.root = b.root.lChild
		} else if data > b.root.data && b.root.rChild != nil {
			b.root = b.root.rChild
		} else {
			break
		}
	}
	return currentElement
}

func isBST(root *Node, min, max int) bool {
	if root == nil {
		return true
	}
	return root.data > min && root.data < max &&
		isBST(root.lChild, min, root.data) &&
		isBST(root.rChild, root.data, max)
}

func (b *BinarySearchTree) printInorderWithConstraints(node *Node, min, max int) {
	if node == nil {
		return
	}
	fmt.Printf("Value: %d, Min: %d, Max: %d\n", node.data, min, max)
	b.printInorderWithConstraints(node.lChild, min, node.data-1)
	b.printInorderWithConstraints(node.rChild, node.data+1, max)
}