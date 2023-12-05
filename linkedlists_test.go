package main

import (
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	if list.head != nil || list.tail != nil {
		t.Errorf("Expected empty list, but head and tail are not nil")
	}
}

func TestAppend(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.Append("apple")
	list.Append("banana")
	if list.head.val != "apple" || list.tail.val != "banana" {
		t.Errorf("Expected head and tail values to be 'apple' and 'banana', got %s and %s", list.head.val, list.tail.val)
	}

	if list.head.next != list.tail || list.tail.prev != list.head {
		t.Errorf("Expected nodes to be linked correctly, but next and prev pointers are incorrect")
	}
}

func TestPrepend(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Prepend(1)
	list.Prepend(2)
	if list.head.val != 2 || list.tail.val != 1 {
		t.Errorf("Expected head and tail values to be 2 and 1, got %d and %d", list.head.val, list.tail.val)
	}
	if list.head.next != list.tail || list.tail.prev != list.head {
		t.Errorf("Expected nodes to be linked correctly, but next and prev pointers are incorrect")
	}
}

func TestGet(t *testing.T) {
	list := NewDoublyLinkedList[float64]()
	list.Append(3.14)
	list.Append(2.71)

	node := list.Get(3.14)
	if node == nil || node.val != 3.14 {
		t.Errorf("Expected to get node with value 3.14, but got nil or %f", node.val)
	}

	node = list.Get(99.99)
	if node != nil {
		t.Errorf("Expected not to find node with value 99.99, but got %f", node.val)
	}
}

func TestGetLength(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.Append("apple")
	list.Append("banana")
	list.Append("orange")

	length := list.GetLength()
	if length != 3 {
		t.Errorf("Expected list length to be 3, got %d", length)
	}
}

func TestRemove(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	removed := list.Remove(2)
	if !removed {
		t.Errorf("Expected to remove node with value 2, but removal failed")
	}

	if list.head.next != list.tail || list.tail.prev != list.head {
		t.Errorf("Expected list structure to remain consistent after removal")
	}

	removed = list.Remove(99)
	if removed {
		t.Errorf("Expected not to remove non-existent element, but removal succeeded")
	}
}

func TestGetAtIndex(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.Append("apple")
	list.Append("banana")
	list.Append("orange")

	node := list.GetAtIndex(1)
	if node == nil || node.val != "banana" {
		t.Errorf("Expected node at index 1 to be 'banana', got %v", node)
	}

	node = list.GetAtIndex(-1)
	if node != nil {
		t.Errorf("Expected nil for index out of bounds (-1), got %v", node)
	}

	node = list.GetAtIndex(list.GetLength() - 1)
	if node == nil || node.val != "orange" {
		t.Errorf("Expected node at last index to be 'orange', got %v", node)
	}
}

func TestContains(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	found, index := list.Contains(2)
	if !found || index != 1 {
		t.Errorf("Expected '2' at index 1, but found %v at %d", found, index)
	}

	found, index = list.Contains(99)
	if found || index != -1 {
		t.Errorf("Expected not to find '99', but found %v at %d", found, index)
	}
}


func TestInsertAt(t *testing.T) {
	list := NewDoublyLinkedList[string]()
	list.Append("apple")
	list.Append("orange")

	insertedNode := list.InsertAt("banana", 1)
	if insertedNode == nil || insertedNode.val != "banana" {
		t.Log("1:")
		t.Errorf("Expected inserted node at index 1 to be 'banana', got %v", insertedNode)
	}

	if list.head.val != "apple" || list.tail.val != "orange" {
		t.Log("1:")
		t.Errorf("Expected original head and tail to remain, but got head: %s, tail: %s", list.head.val, list.tail.val)
	}

	insertedNode = list.InsertAt("kiwi", 0)
	t.Log("1:")
	if insertedNode == nil || insertedNode.val != "kiwi" || list.head != insertedNode {
		t.Errorf("Expected inserted node at index 0 to be 'kiwi' and be new head, got head: %v", list.head)
	}
	list.InsertAt("mango", 4) // Insert at non-existent index should append
	t.Log("1:")
	if list.tail.val != "mango" {
		t.Errorf("Expected inserted node at index 3 to be new tail, but got tail: %s", list.tail.val)
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	removedNode := list.RemoveAt(-1)
	if removedNode != nil {
		t.Errorf("Expected nil for removing out of bounds (-1), got %v", removedNode)
	}
	removedNode = list.RemoveAt(1)
	if removedNode.val != 2 || list.head.val != 1 || list.tail.val != 3 {
		t.Error("error")
	}
	removedNode = list.RemoveAt(0)
	if removedNode == nil || removedNode.val != 1 || list.IsEmpty() {
		t.Errorf("Expected removed node at index 0 to be '1' and leave empty list, got head: %v", list.head)
	}
}
