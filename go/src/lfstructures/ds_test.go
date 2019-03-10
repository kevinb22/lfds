package lfstructures

import (
	"fmt"
	"testing"
)

func TestStackSingleThread(t *testing.T) {
	s := NewLFStack()
	c1 := Container{1}
	c2 := Container{2}
	s.Push(c1)
	s.Push(c2)
	res := s.Pop()
	if res.Get() != c2.Get() {
		t.Fatalf("s.Pop() = %d; want 2", res)
	}
	res = s.Pop()
	if res.Get() != c1.Get() {
		t.Fatalf("s.Pop() = %d; want 1", res)
	}

}

func TestQueueSingleThread(t *testing.T) {
	q := NewLFQueue()
	c1 := Container{1}
	c2 := Container{2}
	q.Produce(c1)
	q.Produce(c2)
	res := q.Consume()
	if res.Get() != c1.Get() {
		t.Fatalf("s.Pop() = %d; want 2", res)
	}
	res = q.Consume()
	if res.Get() != c2.Get() {
		t.Fatalf("s.Pop() = %d; want 2", res)
	}
	fmt.Printf("  ... Passed\n")
}