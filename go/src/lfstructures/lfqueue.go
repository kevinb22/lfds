package lfstructures

import (
	"sync/atomic"
	"unsafe"
)

// LFQueue is implementation of lock free queue in golang
// It is limited by the fact that it is not possible to atomically update both the head and tail
// therefore the base case of the queue has been removed, this queue must start with its Head and Tails
// already formatted as a linked list for example:
// 	testTail := Node{Container{2}, nil}
//	testHead := Node{Container{1}, &testTail}
//	s := LFQueue{&testHead, &testTail}
// 	s.Add(...)
//  s.Remove(...)
type LFQueue struct {
	Head *Node
	Tail *Node
}

// Add appends a value to the Tail of the linked list
func (q *LFQueue) Add(value Container) {
	var oldTail *Node
	newTail := &Node{value, nil}
	for {
		oldTail = q.Tail
		oldTail.Next = newTail
		// atomically update the tail node
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.Tail)),
			unsafe.Pointer(oldTail),
			unsafe.Pointer(newTail)) {
			break
		}
	}
}

//Remove pops from the Head of the linked list
func (q *LFQueue) Remove() Container {
	var oldHead *Node
	var newHead *Node
	for {
		oldHead = q.Head
		if oldHead == nil {
			return nil
		}
		newHead = oldHead.Next
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&q.Head)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(newHead)) {
			break
		}
	}
	return oldHead.Value
}
