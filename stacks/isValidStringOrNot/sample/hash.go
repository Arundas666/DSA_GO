package main
type hashtable struct{
	arr [7]*bucket
}
type bucket struct{
	head *bucketNode
}
type bucketNode struct{
	key int
	nextNode *bucketNode
}
func main(){

}
func hash(key string)int{
	sum:=0
	for _,v:=range key{
		sum+=int(v)

	}
	return sum
}
func (b *bucket)search(key string){
	
}
func (h *hashtable)insert(key string){
	index:=hash(key)
	h.arr[index].insert(key)
}
func(b *bucket) insert(key string){
	
}