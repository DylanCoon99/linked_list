package linkedlist

import (
	"log"
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

func TestPrint(t *testing.T) {
	// tests the printing of the list to stdout



}

func TestAdd(t *testing.T) {

	// first just testing adding one element to an empty list
	var linked_list *LinkedList

	linked_list = InitList()

	linked_list.Add(5)

	if linked_list.Head.Value != 5 || linked_list.Tail.Value != 5 {
		t.Errorf("Failed to add to empty linked list. Expected 5 for head got %d", linked_list.Head.Value)
	}

	log.Println("Successfully added node to empty list.")

	linked_list.Add(3)

	linked_list.Add(7)

	linked_list.PrintList()

}
