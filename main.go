package main

import (
	"AlgoII/Bag"
	"AlgoII/Graph"
	"AlgoII/LinkedList"
	"fmt"
)

func main() {
	//basicTestGraph()
	//testLinkedList()

	graph2 := Graph.BuildGraphFromFile("tinyGraph.txt")
	fmt.Print(graph2.Print())
	startingEdge := 0
	myDSF := Graph.DepthFirstSearch(graph2, startingEdge)
	fmt.Println(myDSF.Print())
	fmt.Println("Path from", startingEdge, "to 3 :\n", myDSF.PathTo(3).Print())

}

func basicTestGraph() {
	bag := Bag.NewBag(2)
	fmt.Print(bag.Print(), " - ", bag.Size(), "\n")
	bag.Add(1, 2, 3, 6, 5, 4)
	fmt.Print(bag.Print(), " - ", bag.Size(), "\n")
	bag.Remove(0, 2)
	fmt.Print(bag.Print(), " - ", bag.Size(), "\n")
	bag.RemoveFirst(4)
	fmt.Print(bag.Print(), " - ", bag.Size(), "\n")

	graph := Graph.NewGraph(5)
	graph.AddEdge(0, 3)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	fmt.Print(graph.Adj(0).Print(), "\n")
	fmt.Print(graph.Adj(2).Print())
}
func testLinkedList() {
	list := LinkedList.NewLinkedList()
	elem1000 := LinkedList.NewListElem(1000)
	for i := 0; i < 2000; i++ {
		ele := LinkedList.NewListElem(i)
		if i == 1000 {
			ele = elem1000
		}
		if i == 1500 {
			list.Push(ele)
			elem1000.SetNext(ele)
		} else {
			list.Push(ele)
		}
	}

	fmt.Println(list.Print())
	fmt.Println("Has loop ? ", list.HasLoop())
}
