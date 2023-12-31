package main

import "fmt"

type graph struct {
	vertices []*vertex
}
type vertex struct {
	data     int
	adjacent []*vertex
}

func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}
func (g *graph) addVertex(data int) {
	if !contains(g.vertices, data) {
		g.vertices = append(g.vertices, &vertex{data: data})
	}
}
func (g *graph) getVertex(data int) *vertex {
	for i, v := range g.vertices {
		if v.data == data {
			return g.vertices[i]
		}
	}
	return nil
}
func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent=append(toVertex.adjacent, fromVertex)

}
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Print("\n vertex ", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}
	}

}
func main() {
	g := graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(5, 3)
	g.addEdge(2, 4)
	g.addEdge(2, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 5)

	g.print()
	fmt.Println(" \n BFS :")
	g.BFS(3)
	fmt.Println(" \n DFS :")
	g.DFS(3)
}

//BFS TRAVERSAL
type queue struct {
	arr []int
}

func (g *graph) BFS(key int) {
	q := queue{}
	var isChecked = make(map[int]bool)
	q.arr=append(q.arr, key)
	isChecked[key] = true
	for len(q.arr) != 0 {
		vertex := q.arr[0]
		q.arr = q.arr[1:]
		fmt.Print(" ", vertex, " ")

		for _, neighbors := range g.getVertex(vertex).adjacent {
			if !isChecked[neighbors.data] {
				isChecked[neighbors.data] = true
				q.arr = append(q.arr, neighbors.data)
			}
		}
		
	}
	fmt.Println(" ")

}
type stack struct{
	arr []int
}
func (g *graph)DFS(value int){
	stack:=stack{}
	var isChecked=make(map[int]bool)
	stack.arr=append(stack.arr, value)
	isChecked[value]=true
	for len(stack.arr)>0{
		vertex:=stack.arr[len(stack.arr)-1]
		stack.arr=stack.arr[:len(stack.arr)-1]
		fmt.Print(" ",vertex," ")
		for _,neighbours:=range g.getVertex(vertex).adjacent{
			if !isChecked[neighbours.data]{
				isChecked[neighbours.data]=true
				stack.arr=append(stack.arr, neighbours.data)
			}
		}

	}

}