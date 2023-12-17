package main

import (
	"fmt"
	"math"
)

type node struct {
	data   int
	left   *node
	right  *node
	parent *node
}
type BinaryTree struct {
	root *node
}

func maxDepth(root *node) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.left)
	rightMax := maxDepth(root.right)
	return int(math.Max(float64(leftMax), float64(rightMax)) + 1)

}
func addNode(root *node, data int) *node {
	if root == nil {
		return &node{data: data}
	}

	if data < root.data {
		root.left = addNode(root.left, data)
	} else {
		root.right = addNode(root.right, data)
	}

	return root
}
func main() {
	b := BinaryTree{}
	b.root=addNode(b.root,6)
	b.root=addNode(b.root,4)
	b.root=addNode(b.root,7)
	b.root=addNode(b.root,8)
	b.root=addNode(b.root,2)
	b.root=addNode(b.root,5)
	maxDepth := maxDepth(b.root)
	fmt.Println("MAx depth is ", maxDepth)

}
