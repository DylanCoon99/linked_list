package linkedlist

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func InitList() *LinkedList {

	return &LinkedList{
		Head: nil,
		Tail: nil,
	}

}

func (l *LinkedList) PrintList() {

	// prints the contents of the list in a pretty way

	output_string := ""

	cur := l.Head

	for cur != nil {
		if cur.Next != nil {
			output_string += (strconv.Itoa(cur.Value) + " -> ")
		} else {
			output_string += (strconv.Itoa(cur.Value))
		}
		cur = cur.Next
	}

	fmt.Println(output_string)
}

func (l *LinkedList) Add(v int) {

	// takes an integer; creates a node; adds it to the end of the list

	node := Node{
		Value: v,
		Next:  nil,
	}

	if l.Head == nil {
		// list is empty
		l.Head = &node
		l.Tail = &node
		return
	}

	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &node
	l.Tail = &node

}



func (l *LinkedList) Insert(val, pos int) {

	// val is the value to insert and pos is the position (0 indexed) to insert it in

	i := 0

	if l.Head == nil {
		l.Add(val)
		return
	}

	node := Node{
		Value: val,
		Next:  nil,
	}

	if pos == 0 {
		// add to the front of the list
		node.Next = l.Head
		l.Head = &node
		return
	}

	cur := l.Head

	for i != pos - 1 {
		cur = cur.Next
		i += 1
	}

	temp := cur.Next

	cur.Next = &node
	node.Next = temp


}