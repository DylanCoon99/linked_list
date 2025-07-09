package linkedlist


import (
	"testing"
)





func TestInit(t *testing.T) {

	var linked_list *LinkedList


	if linked_list = InitList(); linked_list == nil {
		t.Error("Failed to init. Got nil for linked list.")
		return
	}

	if linked_list.Head != nil || linked_list.Tail != nil {
		t.Error("Failed to init. Head and Tail are supposed to be nil.")
	}

}



func TestAdd(t *testing.T) {

	var linked_list *LinkedList

	linked_list = InitList()

	linked_list.Add(5)

}