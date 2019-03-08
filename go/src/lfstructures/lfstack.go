package lfstructures

import (
	"sync/atomic"
	"unsafe"
)

// TrieberStack implementation in golang
type TrieberStack struct {
	Top *Node
}

// Push adds a node onto the top of the TrieberStack
func (ts *TrieberStack) Push(value Container) {
	var oldHead *Node
	newHead := &Node{value, nil}
	for {
		oldHead = ts.Top
		newHead.Next = oldHead
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&ts.Top)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
}

// Pop removes a node from the top of the TrieberStack and returns the node value
func (ts *TrieberStack) Pop() Container {
	var oldHead *Node
	var newHead *Node
	for {
		oldHead = ts.Top
		if oldHead == nil {
			return nil
		}
		newHead = oldHead.Next
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&ts.Top)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
	return oldHead.Value
}
