package main

import "fmt"

type Node struct {
	left  *Node
	key   int
	right *Node
}

type binaryTree struct {
	root *Node
}

func main() {

	node1 := &Node{nil, 7, nil}
	node2 := &Node{nil, 2, nil}
	node3 := &Node{nil, 1, nil}

	node4 := &Node{nil, 0, nil}
	node5 := &Node{nil, 4, nil}

	b := &binaryTree{}

	b.root = node1
	b.root.left = node2
	b.root.right = node3

	b.root.left.left = node4
	b.root.left.right = node5

	fmt.Print("In-order traversal : ")
	inOrder(b.root)
	fmt.Println()

	fmt.Print("Pre-order traversal : ")
	preOrder(b.root)
	fmt.Println()

	fmt.Print("Post-order traversal : ")
	postOrder(b.root)
	fmt.Println()

}

// in order traversal
func inOrder(root *Node) {

	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.key, "  ")
	inOrder(root.right)

}

// pre order traversal
func preOrder(root *Node) {

	if root == nil {
		return
	}

	fmt.Print(root.key, " ")
	preOrder(root.left)
	preOrder(root.right)

}

// post order traversal
func postOrder(root *Node) {

	if root == nil {
		return
	}

	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.key, " ")

}