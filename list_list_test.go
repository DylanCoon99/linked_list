package linkedlist

import (
	"fmt"
	"log"
	"testing"
)

func TestInit(t *testing.T) {

	log.Println("Starting Init Test")

	var linked_list *LinkedList

	if linked_list = InitList(); linked_list == nil {
		t.Error("Failed to init. Got nil for linked list.")
		return
	}

	if linked_list.Head != nil || linked_list.Tail != nil {
		t.Error("Failed to init. Head and Tail are supposed to be nil.")
	}

	log.Println("----- End of Init Test -----")

	fmt.Println()

}

func TestPrint(t *testing.T) {
	// tests the printing of the list to stdout



}




func TestAdd(t *testing.T) {

	log.Println("Starting Add Test")

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

	log.Println("----- End of Add Test -----")

	fmt.Println()
}




func TestInsert(t *testing.T) {

	// Test the functionality of insert to different indices in the list
	log.Println("Starting Insert Test")

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

	log.Println("----- End of Insert Test -----")

	fmt.Println()
}


func TestSearch(t *testing.T) {

	log.Println("Starting Search Test")

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

	log.Println("----- End of Search Test -----")

	fmt.Println()
}



func TestDelete(t *testing.T) {

	// Tests to make sure the delete functionality works
	log.Println("Starting Delete Test")

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


	if ret := linked_list.Delete(88); !ret {
		t.Error("Failed to delete 88 from linked list.")
	}

	// Expected State: 11 -> 5 -> 7 -> 9 -> 2 -> 99
	linked_list.PrintList()

	if ret := linked_list.Delete(99); !ret {
		t.Error("Failed to delete 99 from linked list.")
	}

	// Expected State: 11 -> 5 -> 7 -> 9 -> 2
	linked_list.PrintList()

	if ret := linked_list.Delete(88); ret {
		t.Error("Failed to return false when 88 was not present in list.")
	}

	// Expected State: 11 -> 5 -> 7 -> 9 -> 2
	linked_list.PrintList()

	log.Println("----- End of Delete Test -----")
	fmt.Println()

}