package lfstructures

import (
	"fmt"
	"testing"
)

func TestSingleThread(t *testing.T) {
	ts := TrieberStack{&Node{1, nil}}
	ts.Push(2)
	res := ts.Pop()
	if res != 2 {
		t.Fatalf("[ts.Push(2), ts.Pop()] = %d; want 2", res)
	}
	fmt.Printf("  ... Passed\n")
}
