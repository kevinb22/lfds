package lfstructures

import (
	"sync/atomic"
	"unsafe"
)

// LFStack is a TrieberStack implementation in golang
type LFStack struct {
	Top *Node
}

// Push adds a node onto the Top of the TrieberStack
func (s *LFStack) Push(value Container) {
	var oldHead *Node
	newHead := &Node{value, nil}
	for {
		oldHead = s.Top
		newHead.Next = oldHead
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
}

// Pop removes a node from the Top of the TrieberStack and returns the node value
func (s *LFStack) Pop() Container {
	var oldHead *Node
	var newHead *Node
	for {
		oldHead = s.Top
		if oldHead == nil {
			return nil
		}
		newHead = oldHead.Next
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
	return oldHead.Value
}
