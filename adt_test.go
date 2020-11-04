package main

import (
	"testing"
)

func TestHeap(t *testing.T){

	n1 := TNode{2, 5}
	n2 := TNode{3, 2}
	n3 := TNode{4, 1}
	n4 := TNode{4, 6}
	n5 := TNode{5, 5}

	heap := PriorityQueue{}
	sample := []TNode{n1, n2, n3, n4, n5}
	heap.arr = append(heap.arr, sample...)
	heap.Insert(TNode{2, 10})

	// fmt.Println("heap", heap)
	// fmt.Println("min", heap.ExtractMin()) // {4, 1}
	// fmt.Println("heap", heap)
	// fmt.Println("min", heap.ExtractMin()) // {3, 2}
	// fmt.Println("heap", heap)
	// fmt.Println("min", heap.ExtractMin()) // {5, 5}
	// fmt.Println("heap", heap)

	min1 := heap.ExtractMin()
	if min1.w != 1 {
		t.Errorf("min1 = %d; want 1", int(min1.w))
	}

	min2 := heap.ExtractMin()
	if min2.w != 2 {
		t.Errorf("min2 = %d; want 2", int(min2.w))
	}

	min3 := heap.ExtractMin()
	if min3.w != 5 {
		t.Errorf("min3 = %d; want 5", int(min3.w))
	}
}