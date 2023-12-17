package main
type minHeap struct{
	arr []int
}
func(h *minHeap)build(arr []int){
	h.arr=arr
	for i:=parent(len(arr)-1);i>0;i--{
		h.shiftDown(i)
	}

}
func (h *minHeap)shiftDown(currentIndex int){
		leftIndex:=leftChild(currentIndex)
		rightIndex:=rightChild(currentIndex)
		endIndex:=len(h.arr)-1
		var indxToSwap int
		for leftIndex<=endIndex{
			if rightIndex<=endIndex && h.arr[rightIndex]<h.arr[leftIndex]{
				indxToSwap=rightIndex
			}else{
				indxToSwap=leftIndex
			}
			if h.arr[indxToSwap]<h.arr[currentIndex]{
				swap(h.arr,indxToSwap,currentIndex)
				currentIndex=indxToSwap
				leftIndex=leftChild(currentIndex)
			}
		}
}
func parent(i int)int{
	return (i-1)/2
}
func leftChild(i int)int{
	return (2*i)+1
}
func rightChild(i int)int{
	return (2*i)+2
}
func swap(arr []int,a int, b int){
	arr[a],arr[b]=arr[b],arr[a]
}
func main(){

}