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


func (l *LinkedList) SearchList(val int) int {

	// returns the index of the first occurence of a value 
	// -1 if not present

	i := 0

	cur := l.Head

	for cur != nil {
		if cur.Value == val {
			return i
		}
		i += 1
		cur = cur.Next
	}

	return -1
}



func (l *LinkedList) Delete(val int) bool {

	// deletes the first occurence of the val in the list
	// return true is the element was found and deleted; false otherwise

	if l.Head == nil {
		return false
	}

	if l.Head.Value == val {
		l.Head = nil
		l.Tail = nil
		return true
	}


	cur := l.Head

	for cur.Next != nil {

		if cur.Next.Value == val {
			cur.Next = cur.Next.Next 
			return true
		}

		cur = cur.Next
	}


	return false

}



func (l *LinkedList) Reverse() {

	// reverses the linked list in places

	if l.Head == nil || l.Head.Next == nil {
		// if the list is empty or is only 1 node; just return
		return 
	}

	var prev *Node
	
	prev = nil
	cur := l.Head

	l.Tail = l.Head

	next := cur.Next

	for cur.Next != nil {
		// cur points to prev; cur becomes next; prev becomes cur
		cur.Next = prev
		prev = cur
		cur = next
		next = cur.Next
	}

	cur.Next = prev

	l.Head = cur

	return

}