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


/*

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


*/

func TestInsert(t *testing.T) {

	// Test the functionality of insert to different indices in the list


	// Init the list
	var linked_list *LinkedList

	linked_list = InitList()

	// Add some values to the list
	linked_list.Add(5)
	linked_list.Add(7)
	linked_list.Add(9)
	linked_list.Add(2)

	linked_list.PrintList()
	// Current State: 5 -> 7 -> 9 -> 2

	// Insert into the front of the list
	linked_list.Insert(11, 0) 

	// Expected State: 11 -> 5 -> 7 -> 9 -> 2
	linked_list.PrintList()

	// Insert into the middle of the list
	linked_list.Insert(88, 2)

	// Expected State: 11 -> 5 -> 88 -> 7 -> 9 -> 2
	linked_list.PrintList()


	// Insert at the end of the list
	linked_list.Insert(99, 6)

	// Expected State: 11 -> 5 -> 88 -> 7 -> 9 -> 2 -> 99
	linked_list.PrintList()

}


func TestSearch(t *testing.T) {

	// Init the list
	var linked_list *LinkedList

	linked_list = InitList()

	// Add some values to the list
	linked_list.Add(5)
	linked_list.Add(7)
	linked_list.Add(9)
	linked_list.Add(2)

	// Insert into the front of the list
	linked_list.Insert(11, 0) 

	// Insert into the middle of the list
	linked_list.Insert(88, 2)

	// Insert at the end of the list
	linked_list.Insert(99, 6)

	// Expected State: 11 -> 5 -> 88 -> 7 -> 9 -> 2 -> 99
	linked_list.PrintList()

	r := linked_list.SearchList(88) // Expected Result: 2

	if r != 2 {
		t.Errorf("Failed to search linked list. Expected 2 got %d", r)
	}


	r = linked_list.SearchList(100) // Expected Result: 2

	if r != -1 {
		t.Errorf("Failed to search linked list. Expected -1 got %d", r)
	}

	log.Println("Successfully searched the linked list.")

}