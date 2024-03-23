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
	val      string
	nextNode *bucketNode
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)

	}
	x := sum % arraySize
	return x
}
func (h *hashtable) Insert(key, val string) {
	index := hash(key)
	_, ok := h.Search(key)

	if !ok {
	h.arr[index].insert(key, val)
	} else {
		fmt.Println("already exist")
	}
}
func (b *bucket) insert(key, val string) {
		newNode := &bucketNode{key: key, val: val}
		newNode.nextNode = b.head
		b.head = newNode
}

func (h1 *hashtable) Search(key string) (string, bool) {

	index := hash(key)

	return h1.arr[index].searchBucket(key)
}

func (b1 *bucket) searchBucket(key string) (string, bool) {

	current := b1.head
	if current == nil {
		return "", false
	}

	for current!= nil {
		if current.key == key {
			return current.val, true
		}
		current = current.nextNode
	}

	return "", false

}

func main() {
	hash := &hashtable{}
	for i := range hash.arr {
		hash.arr[i] = &bucket{}
	}
	hash.Insert("ARUN", "GOOD BIY")
	hash.Insert("akhil", "bad BIYaa")
	
	x, ok := hash.Search("akhil")
	fmt.Println(x, ok)

}
