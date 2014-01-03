package Bag

import (
	"fmt"
	"strconv"
)

type Bag interface {
	Add(...int)
	Remove(int, int)
	RemoveFirst(int)
	Size() int
	Print() string
	Elements() []int
}
type MyBag struct {
	elem *[]int
}

func NewBag(size int) Bag {
	bag := new(MyBag)
	newElem := make([]int, size)
	bag.elem = &newElem
	return bag
}
func NewEmptyBag() Bag {
	bag := new(MyBag)
	newElem := []int{}
	bag.elem = &newElem
	return bag
}
func (b MyBag) Size() int {
	return len(*b.elem)
}
func (b MyBag) Elements() []int {
	return (*b.elem)
}
func (b *MyBag) Add(elems ...int) {
	if b.elem == nil {
		b.elem = &elems
		return
	}
	*b.elem = append(*b.elem, elems...)
}
func (b *MyBag) Remove(elem int, rank int) {
	if b.elem == nil {
		panic(fmt.Sprintf("[Remove] Error empty bad"))
	}
	if len(*b.elem) == 1 {
		b.elem = nil
	}
	CurrRank := 0
	shiftIndex := 0
	newElem := make([]int, len(*b.elem))
	for i, e := range *b.elem {
		if e == elem {
			CurrRank++
			if rank == CurrRank {
				shiftIndex = 1 // If one elem is removed we shift by one
				continue
			}
		}
		newElem[i-shiftIndex] = (*b.elem)[i]
	}
	*b.elem = newElem[:len(*b.elem)-shiftIndex]
}
func (b *MyBag) RemoveFirst(elem int) {
	b.Remove(elem, 1)
}
func (b MyBag) Print() string {
	if b.elem == nil {
		return "Empty bag"
	}
	result := ""
	for _, e := range *b.elem {
		result += strconv.Itoa(e) + " "
	}
	return result
}
