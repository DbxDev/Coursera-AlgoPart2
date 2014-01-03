package Graph

import (
	"AlgoII/LinkedList"
	"fmt"
	"strconv"
)

const NOT_SET int = -1

type DFS interface {
	HasPathTo(int) bool
	PathTo(int) LinkedList.LinkedList
	Marked() []bool
	EdgeTo() []int
	StartingEdge() int
	Graph() Graph
	Print() string
}

type MyDFS struct {
	graph        Graph
	startingEdge int
	marked       []bool
	edgeTo       []int
}

func (dfs MyDFS) Marked() []bool {
	return dfs.marked
}
func (dfs MyDFS) EdgeTo() []int {
	return dfs.edgeTo
}
func (dfs MyDFS) StartingEdge() int {
	return dfs.startingEdge
}
func (dfs MyDFS) Graph() Graph {
	return dfs.graph
}

func DepthFirstSearch(g Graph, s int) DFS {
	theDfs := new(MyDFS)
	theDfs.graph = g
	theDfs.startingEdge = s
	theDfs.marked = make([]bool, theDfs.graph.V())
	theDfs.edgeTo = make([]int, theDfs.graph.V())
	for i := range theDfs.edgeTo {
		theDfs.edgeTo[i] = NOT_SET
	}
	// start dfs
	dfs(s, theDfs)

	return theDfs
}
func dfs(edge int, theDfs DFS) {
	theDfs.Marked()[edge] = true
	for _, v := range theDfs.Graph().Adj(edge).Elements() {
		if !theDfs.Marked()[v] {
			theDfs.EdgeTo()[v] = edge
			dfs(v, theDfs)
		}
	}
}

func (dfs MyDFS) HasPathTo(v int) bool {
	return dfs.Marked()[v]
}
func (dfs MyDFS) PathTo(v int) LinkedList.LinkedList {
	if !dfs.HasPathTo(v) {
		fmt.Printf("Edge [%d] is not connected to [%d]\n", v, dfs.StartingEdge())
		return nil
	}
	path := LinkedList.NewLinkedList()
	edge := dfs.EdgeTo()[v]
	path.Push(LinkedList.NewListElem(v))
	for edge != dfs.StartingEdge() && edge != NOT_SET {
		edgeElem := LinkedList.NewListElem(edge)
		path.Push(edgeElem)
		edge = dfs.EdgeTo()[edge]
	}
	return path
}
func (dfs MyDFS) Print() string {
	result := "DFS on graph : " + dfs.graph.Print() + "Starting at [" + strconv.Itoa(dfs.startingEdge) + "]\n"
	result += "Marked array : \n"
	for i, v := range dfs.marked {
		result += "[" + strconv.Itoa(i) + "] : " + strconv.FormatBool(v) + "\n"
	}

	result += "EdgeTo array : \n"
	for i, v := range dfs.edgeTo {
		result += "[" + strconv.Itoa(i) + "] : " + strconv.Itoa(v) + "\n"
	}
	return result
}
