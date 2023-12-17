package main

import "fmt"

type BinaryTree struct {
	root *node
}
type node struct {
	data   int
	left   *node
	right  *node
	parent *node
}

func addNode(root *node, data int) *node {
	newNode := &node{data: data}
	if root == nil {
		root = newNode
		return root
	}
	if root.data < data {
		root.right = addNode(root.right, data)
	} else {
		root.left = addNode(root.left, data)
	}
	return root
}
func (b *BinaryTree) SymmetricOrNot() bool {
	if b.root == nil {
		return true
	}
	return isSymmetric( b.root.left, b.root.right)

}

func isSymmetric(left *node, right *node) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }

    return isSymmetric(left.left, right.right) && isSymmetric(left.right, right.left)
}

func main() {
	b := BinaryTree{}
	b.root = addNode(b.root, 6)
	b.root = addNode(b.root, 4)
	b.root = addNode(b.root, 8)
	b.root = addNode(b.root, 5)
	b.root = addNode(b.root, 3)
	b.root = addNode(b.root, 9)
	b.root = addNode(b.root, 7)
	

	sYmm := b.SymmetricOrNot()
	fmt.Println(sYmm)

}
