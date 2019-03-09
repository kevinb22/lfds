package lfstructures

import (
	"fmt"
	"testing"
)

func TestStackSingleThread(t *testing.T) {
	s := LFStack{}
	s.Push(Container{1})
	s.Push(Container{2})
	res := s.Pop()
	if res.Get() != 2 {
		t.Fatalf("s.Pop() = %d; want 2", res)
	}
	res = s.Pop()
	if res.Get() != 1 {
		t.Fatalf("s.Pop() = %d; want 1", res)
	}
	fmt.Printf("  ... Passed\n")
}

func TestQueueSingleThread(t *testing.T) {
	testTail := Node{Container{2}, nil}
	testHead := Node{Container{1}, &testTail}
	q := LFQueue{&testHead, &testTail}
	q.Add(Container{3})
	res := q.Remove()
	if res.Get() != 1 {
		t.Fatalf("s.Pop() = %d; want 1", res)
	}
	res = q.Remove()
	if res.Get() != 2 {
		t.Fatalf("s.Pop() = %d; want 2", res)
	}
	res = q.Remove()
	if res.Get() != 3 {
		t.Fatalf("s.Pop() = %d; want 3", res)
	}
	fmt.Printf("  ... Passed\n")
}
