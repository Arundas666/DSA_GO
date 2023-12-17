package main
import "fmt"
func main() {
	var arr = []int{4, 3, 2, 1, 5}
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
