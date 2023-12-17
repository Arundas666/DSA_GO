package main

import "fmt"

const size = 26

type node struct {
	children [size]*node
	isEnd    bool
}
type trie struct {
	root *node
}
func (t *trie) insert(w string) {
	currentIdx := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentIdx.children[charIndex] == nil {
			currentIdx.children[charIndex] = &node{}
		}
		currentIdx = currentIdx.children[charIndex]
	}
	currentIdx.isEnd = true
}
func (t *trie) search(w string) bool {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd {
		return true
	}
	return false
}
func main() {
	t := trie{&node{}}
	t.insert("arundas")
	t.insert("arunda")
	b := t.search("arunda")
	fmt.Println(b)

}
