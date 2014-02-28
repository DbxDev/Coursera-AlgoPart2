package UnionFind

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const valueSeparator string = " "

type UF interface {
	Union(int, int)
	Find(int) int
	Connected(int, int) bool
	Size() int
}

type MyUF struct {
	id   []int
	size int
}

func NewUF(N int) UF {
	uf := new(MyUF)
	uf.id = make([]int, N)
	uf.size = N
	for i := 0; i < N; i++ {
		uf.id[i] = i
	}
	return uf
}
func (uf *MyUF) Union(p int, q int) {
	qRoot := uf.Find(q)
	pRoot := uf.Find(p)
	if pRoot != qRoot {
		uf.id[p] = q
	}
}

func (uf MyUF) Find(p int) int {
	root := p
	for uf.id[root] != root {
		root = uf.id[root]
	}
	return root
}

func (uf MyUF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}
func (uf MyUF) Size() int {
	return uf.size
}
func BuildUFFromFile(inputFileName string) UF {
	var uf UF
	data, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(fmt.Sprintf("Error while reading file %v. Error : %v", inputFileName, err))
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
			size, _ := strconv.Atoi(scannerLine.Text())
			uf = NewUF(size)
		} else {
			values := strings.Split(scannerLine.Text(), valueSeparator)
			v, _ := strconv.Atoi(values[0])
			w, _ := strconv.Atoi(values[1])
			uf.Union(v, w)
		}
	}

	return uf
}
func Print(uf UF) string {
	root := make([]string, uf.Size())
	result := ""
	for i, _ := range root {
		root[uf.Find(i)] += strconv.Itoa(i) + "-"
	}
	for _, elem := range root {
		if elem != "" {
			result += elem + "\n"
		}
	}
	return result
}
