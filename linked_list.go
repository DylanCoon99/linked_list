package linkedlist


type Node struct {
	Value int
	Next *Node
}


type LinkedList struct {
	Head *Node
	Tail *Node
}



func InitList() *LinkedList{


	return &LinkedList{
		Head: nil,
		Tail: nil,
	}


}



func PrintList(l *LinkedList) {

	// prints the contents of the list in a pretty way

}



func (l *LinkedList) Add(v int) {

	// takes an integer; creates a node; adds it to the end of the list

	node := Node {
		Value: v,
		Next: nil,
	}


	if l.head == nil {
		// list is empty
		l.head = node
		l.tail = node
		return
	}




}





