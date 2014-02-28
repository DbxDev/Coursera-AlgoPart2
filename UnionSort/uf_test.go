package UnionFind

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	uf := NewUF(5)
	const node1, node2 = 1, 2
	uf.Union(node1, node2)
	result := uf.Connected(node1, node2)
	if !result {
		t.Errorf("Connected(%v,%v) = %v, want %v", node1, node2, result, true)
	}
}

// 3-4-8-9 and 0-1-2-5-6-7
func TestBuildFromSmallFile(t *testing.T) {
	const path = "../TestFiles/tinyUF.txt"
	const size = 10
	uf := BuildUFFromFile(path)
	result := uf.Size() == size
	if !result {
		t.Errorf("Building UF from file of size %v. Expected %v", uf.Size(), size)
	}
	var node1, node2 = 4, 9
	result = uf.Connected(node1, node2)
	if !result {
		t.Errorf("Connected(%v , %v) = %v. Expected %v", node1, node2, result, true)
	}
	node1, node2 = 0, 9
	result = uf.Connected(node1, node2)
	if result {
		t.Errorf("Connected(%v , %v) = %v. Expected %v", node1, node2, result, false)
	}
	fmt.Println(Print(uf))
}
func TestBuildFromSmallFile(t *testing.T) {
	const path = "../TestFiles/tinyUF.txt"
	const size = 625
	uf := BuildUFFromFile(path)
	result := uf.Size() == size
	if !result {
		t.Errorf("Building UF from file of size %v. Expected %v", uf.Size(), size)
	}
	var node1, node2 = 4, 9
	result = uf.Connected(node1, node2)
	node1, node2 = 0, 9
	result = uf.Connected(node1, node2)

	//fmt.Println(Print(uf))
}
