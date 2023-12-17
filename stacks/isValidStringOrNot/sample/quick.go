package main

import "fmt"

func main() {
	var arr = []int{4, 5, 3, 7, 2, 6}
	sort(arr,0,5)
	fmt.Println(arr)
}
func sort(arr []int,start int,end int)[]int{
	if start>=end{
		return nil
	}
	pivot:=start
	left:=start+1
	right:=end
	for left<=right{
		if arr[left]<arr[pivot]&&arr[right]>arr[pivot]{
			arr[left],arr[right]=arr[right],arr[left]
			left++
			right--
		}
		if arr[left]>=arr[pivot]{
			left++
		}
		if arr[right]<=arr[pivot]{
			right--
		}
	}
	arr[right],arr[pivot]=arr[pivot],arr[right]
	sort(arr,start,right-1)
	sort(arr,right+1,end)
	return arr


}
