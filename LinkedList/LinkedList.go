package LinkedList

import (
	"fmt"
	"strconv"
)

type LinkedList interface {
	Push(ListElement)
	Size() int
	Print() string
	HasLoop() bool
}
type ListElement interface {
	Value() int
	Next() ListElement
	SetNext(ListElement)
}

type MyLinkedList struct {
	elem ListElement // first
}
type MyListElem struct {
	value int
	next  ListElement
}

func (e MyListElem) Value() int {
	return e.value
}
func (e MyListElem) Next() ListElement {
	return e.next
}
func (e *MyListElem) SetNext(next ListElement) {
	e.next = next
}
func NewListElem(v int) ListElement {
	elem := new(MyListElem)
	elem.value = v
	return elem
}
func NewLinkedList() LinkedList {
	return new(MyLinkedList)
}
func (e *MyLinkedList) Push(ele ListElement) {
	if e.elem == nil {
		e.elem = ele
		return
	}
	ele.SetNext(e.elem)
	e.elem = ele
}

func (e *MyLinkedList) Size() int {
	if e.elem == nil {
		return 0
	}
	if e.HasLoop() {
		panic(fmt.Sprintf("The list has loop. Cannot compute size."))
	}
	elem := e.elem
	count := 1
	for elem.Next() != nil {
		elem = elem.Next()
		count++
	}
	return count
}
func (e *MyLinkedList) Print() string {
	if e.elem == nil {
		return "Empty list"
	}
	if e.HasLoop() {
		return "The list has loop. Cannot print"
	}
	result := "List :\n"
	elem := e.elem
	for elem.Next() != nil {
		result += strconv.Itoa(elem.Value()) + " -> " + strconv.Itoa(elem.Next().Value()) + "\n"
		elem = elem.Next()
	}
	result += strconv.Itoa(elem.Value()) + " has no next elem\n"
	return result
}
func (e *MyLinkedList) HasLoop() bool {
	if e.elem == nil {
		return false
	} else if e.elem.Next() == nil {
		return false
	}
	elemSlow := e.elem
	elemFast := e.elem.Next()
	if elemSlow == elemFast {
		return true
	}
	for {
		elemSlow = elemSlow.Next()
		// end of list > No loop
		if elemFast.Next() != nil && elemFast.Next().Next() != nil {
			elemFast = elemFast.Next().Next()
		} else {
			return false
		}
		// Found elem[i] = elem[ 2 * i ] in the loop
		if elemFast == elemSlow {
			return true
		}

	}
	return false
}
