package Graph

import (
	"AlgoII/Bag"
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const valueSeparator string = " "

type Graph interface {
	AddEdge(int, int)
	Adj(int) Bag.Bag
	Print() string
	V() int
}
type MyGraph struct {
	v   int
	adj *[]Bag.Bag
}

func NewGraph(V int) Graph {
	g := new(MyGraph)
	newGraph := make([]Bag.Bag, V)
	for i := range newGraph {
		newGraph[i] = Bag.NewEmptyBag()
	}
	g.v = V
	g.adj = &newGraph
	return g
}
func BuildGraphFromFile(inputFileName string) Graph {
	var graph Graph
	data, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(fmt.Sprint("Error while reading file %s", inputFileName))
		return nil
	}
	// a scanner line by line
	scannerLine := bufio.NewScanner(strings.NewReader(string(data)))
	count := 0
	for scannerLine.Scan() {
		count++
		if err = scannerLine.Err(); err != nil {
			break
		}
		// Line 1 : Number of vertices
		if count == 1 {
			vectives, _ := strconv.Atoi(scannerLine.Text())
			graph = NewGraph(vectives)
		} else if count == 2 {
			//edges, _ := strconv.Atoi(scannerLine.Text())
		} else {
			values := strings.Split(scannerLine.Text(), valueSeparator)
			v, _ := strconv.Atoi(values[0])
			w, _ := strconv.Atoi(values[1])
			graph.AddEdge(v, w)
		}
	}

	return graph
}

// Add and edge W to V
func (g *MyGraph) AddEdge(v int, w int) {
	(*g.adj)[v].Add(w)
	(*g.adj)[w].Add(v)
}
func (g MyGraph) V() int {
	return g.v
}
func (g MyGraph) Adj(v int) Bag.Bag {
	return (*g.adj)[v]
}
func (g MyGraph) Print() string {
	result := "Graph : \n"
	for i, edge := range *g.adj {
		result += "[" + strconv.Itoa(i) + "] - " + edge.Print() + "\n"
	}
	return result
}
