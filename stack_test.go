package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := NewStack[int]()

	// Push elements onto the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Check if the stack length is updated correctly
	if stack.GetLength() != 3 {
		t.Error("Stack length should be 3, but got", stack.GetLength())
	}

	// Check if the stack top value is updated correctly
	if stack.Peek() != 3 {
		t.Error("Stack top value should be 3, but got", stack.Peek())
	}
}

func TestStackIsEmpty(t *testing.T) {
	stack := NewStack[int]()

	// Check if the empty stack is empty
	if !stack.IsEmpty() {
		t.Error("Empty stack should be empty")
	}

	// Push an element onto the stack
	stack.Add(1)

	// Check if the non-empty stack is not empty
	if stack.IsEmpty() {
		t.Error("Non-empty stack should not be empty")
	}
}

func TestStackGetLength(t *testing.T) {
	stack := NewStack[int]()

	// Check the length of the empty stack
	if stack.GetLength() != 0 {
		t.Error("Empty stack length should be 0, but got", stack.GetLength())
	}

	// Push elements onto the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Check the length of the non-empty stack
	if stack.GetLength() != 3 {
		t.Error("Stack length should be 3, but got", stack.GetLength())
	}
}

func TestStackContains(t *testing.T) {
	stack := NewStack[int]()

	// Add elements to the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Check if the stack contains element 2
	contains, index := stack.Contains(2)
	if !contains || index != 1 {
		t.Error("Stack should contain element 2 at index 1")
	}

	// Check if the stack does not contain element 4
	contains, index = stack.Contains(4)
	if contains || index != -1 {
		t.Error("Stack should not contain element 4")
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStack[int]()

	// Add elements onto the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Pop an element from the stack
	poppedNode := stack.Pop()

	// Check if the popped element is correct
	if poppedNode.val != 3 {
		t.Error("Popped element value should be 3, but got", poppedNode.val)
	}

	// Check if the stack length is updated correctly
	if stack.GetLength() != 2 {
		t.Error("Stack length should be 2, but got", stack.GetLength())
	}
}

func TestStackPeek(t *testing.T) {
	stack := NewStack[int]()

	// Add elements to the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Peek at the top value
	peekedValue := stack.Peek()

	// Check if the peeked value is correct
	if peekedValue != 3 {
		t.Error("Stack peeked value should be 3, but got", peekedValue)
	}

	// Check if the stack remains unchanged after peeking
	if stack.GetLength() != 3 {
		t.Error("Stack length should remain unchanged after peeking")
	}
}

func TestStackPrintData(t *testing.T) {
	stack := NewStack[int]()

	// Add elements to the stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	t.Log("Stack contents:")
	// Print the contents of the stack
	stack.PrintData()

	// The expected output should be:
	// 3 -> 2 -> 1
	t.Log("Stack test passed")
}
