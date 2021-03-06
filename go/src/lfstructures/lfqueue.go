package lfstructures

import (
	"sync/atomic"
	"unsafe"
)

// LFQueue is implementation of a one producer, one consumer lock free queue in golang
// The producer owns all nodes before divider, the next pointer inside the last node,
// and the ability to update first and last. The consumer owns everything else, including
// the values in the nodes from divider onward, and the ability to update divider.
type LFQueue struct {
	First, Divider, Last *Node // First is for produce user only, Divider and Last can be shared
	// Divider and Last need to be treated as atomic variable
}

func NewLFQueue() *LFQueue {
	q := new(LFQueue)
	q.First = &Node{nil, nil}
	q.Divider = q.First
	q.Last = q.First
	return q
}

// Add appends a value to the Tail of the linked list
func (q *LFQueue) Produce(value Container) {
	// Add new item
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&q.Last.Next)),
		unsafe.Pointer(&Node{value, nil}))
	// Publish item
	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&q.Last)),
		unsafe.Pointer(q.Last),
		unsafe.Pointer(q.Last.Next))
	// Trim unused nodes
	for q.First != q.Divider {
		q.First = q.First.Next
	}
}

// Remove pops from the Head of the linked list
func (q *LFQueue) Consume() Container {
	// If queue is not empty
	oldDivider := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&q.Divider)))
	oldLast := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&q.Last)))
	if oldDivider != oldLast {
		// Grab the next consumable result
		result := atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&q.Divider.Next)))
		// Publish this update
		atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&q.Divider)),
			unsafe.Pointer(q.Divider),
			unsafe.Pointer(q.Divider.Next))
		// Return the payload of the result
		return (*Node)(result).Value
	}
	return nil
}
