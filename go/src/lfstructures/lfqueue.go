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
	atomicLock, locked, unlocked int32
}

func NewLFQueue() *LFQueue {
	q := new(LFQueue)
	q.First = &Node{nil, nil}
	q.Divider = q.First
	q.Last = q.First
	q.atomicLock = int32(0)
	q.locked = int32(1)
	q.unlocked = int32(0)
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
// Not locking introduces a bug I was unable to figure out in time.
// So I admitted defeat and must use a spinlock for now.
func (q *LFQueue) Consume() Container {
	q.lock()
	defer q.unlock()
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

// Code below adopted from https://medium.com/@tylerneely/fear-and-loathing-in-lock-free-programming-7158b1cdd50c
func (q *LFQueue) lock() {
	pointer := &q.atomicLock
	old := q.unlocked
	new := q.locked
	for {
		// spin until we successfully change the
		// atomicLock from unlocked to locked
		if atomic.CompareAndSwapInt32(pointer, old, new) {
			return
		}
	}
}

func (q *LFQueue) unlock() {
	atomic.StoreInt32(&q.atomicLock, q.unlocked)
}
