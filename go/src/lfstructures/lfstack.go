package lfstructures

import (
	"sync/atomic"
	"unsafe"
)

// LFStack is a TrieberStack implementation in golang
type LFStack struct {
	Top *Node // Needs to be treated as Atomic reference
}

// NewLFStack returns a pointer to a new LFStack initialized on the heap
func NewLFStack() *LFStack {
	s := new(LFStack)
	s.Top = nil
	return s
}

// Push adds a node onto the Top of the TrieberStack
func (s *LFStack) Push(value Container) {
	var oldHead *Node
	newHead := &Node{value, nil}
	for {
		// Atomically load current head (oldHead = s.Top)
		oldHead = (*Node) (atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top))))
		// Make local update for newHead
		newHead.Next = oldHead
		// If CAS works then break, else try again
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
		// Atomically load current head, (oldHead = s.Top)
		oldHead = (*Node) (atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top))))
		// If head null stack is empty
		if oldHead == nil {
			return nil
		}
		// Atomically load new head, (newHead = oldHead.Next)
		newHead = (*Node) (atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top.Next))))
		// If CAS works then break, else try again
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.Top)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
	return oldHead.Value
}
