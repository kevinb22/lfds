package lfstructures

import (
	"fmt"
	"sync"
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

func threadStackRoutine(s *LFStack, wg *sync.WaitGroup) {
	c1 := Container{1}
	s.Push(c1)
	s.Pop()
	wg.Done()
}

func TestStackMultiThread(t *testing.T) {
	s := NewLFStack()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go threadStackRoutine(s, &wg)
	}
	wg.Wait()
	if s.Top != nil {
		t.Fatalf("s.Top = %v; want nil", s.Top)
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

func threadQueueProduceRoutine(s *LFQueue, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		c1 := Container{1}
		s.Produce(c1)
	}

	wg.Done()
}
func threadQueueConsumeRoutine(s *LFQueue, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		s.Consume()
	}
	wg.Done()
}

func TestQueueMultiThread(t *testing.T) {
	q := NewLFQueue()
	var wg sync.WaitGroup
	wg.Add(2)
	go threadQueueProduceRoutine(q, &wg)
	go threadQueueConsumeRoutine(q, &wg)
	wg.Wait()
	res := (q.First == q.Divider) || (q.Divider == q.Last)
	if !res {
		t.Fatalf("q.First = %v; q.Divider = %v; q.Last = %v; two of the three should be equal", q.First, q.Divider, q.Last)
	}
	fmt.Printf("  ... Passed\n")
}
