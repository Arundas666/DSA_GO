package main

import "fmt"

const arraySize = 7

type hashtable struct {
	arr [arraySize]*bucket
}
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key      string
	nextNode *bucketNode
}

func main() {
	hash := &hashtable{}
	for i:=range hash.arr{
		hash.arr[i]=&bucket{}
	}
	hash.Insert("HAI")
	hash.Insert("HEY")
	hash.Insert("HELLO")
	hash.Insert("HOOI")
	fmt.Println(hash)
	hash.Delete("HELLO")
	fmt.Println(hash.arr)
}
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)

	}
	return sum % arraySize
}
func (h *hashtable) Insert(key string) {
	index := hash(key)
	h.arr[index].insert(key)
}
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.nextNode = b.head
		b.head = newNode
	} else {
		fmt.Println("already exist")
	}

}
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.nextNode
	}
	return false
}
func (h *hashtable) Delete(key string) {
	index := hash(key)
	h.arr[index].delete(key)

}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.nextNode
		return
	}
	prevNode := b.head
	for prevNode.nextNode != nil {
		if prevNode.nextNode.key == key {
			prevNode.nextNode = prevNode.nextNode.nextNode
		}
		prevNode = prevNode.nextNode
	}
}
