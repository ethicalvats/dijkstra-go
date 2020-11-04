package main

// Implements a minHeap 
type PriorityQueue struct {
	arr []TNode
}

func (*PriorityQueue) LeftChild(i int) int{
	return 2 * i + 1
}

func (*PriorityQueue) RightChild(i int) int{
	return 2 * i + 2
}

func (p *PriorityQueue) swap(i int, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}

// extract min element (root), swap root node with last element, run minHeapify
func (p *PriorityQueue) ExtractMin() TNode {
	min := p.arr[0]
	size := len(p.arr)
	p.swap(0, size - 1)
	p.arr = p.arr[:size-1]
	p.MinHeapify()
	return min
}

// insert a new node at root, run minHeapify
func (p *PriorityQueue) Insert(a TNode){
	p.arr = append(p.arr, a)
	p.MinHeapify()
}

//finds and updates the weights
func (p *PriorityQueue) Update(v int, w float64){
	for _, k := range p.arr {
		if k.v == v {
			k.w = w
		}
	}
}

// compare parent node with its leaves recursively and swap in invariant (parent is smaller than childrens) not satisfied
func (p *PriorityQueue) MinHeapify() {
	size := len(p.arr)
	if size > 1 {
		for i := size/2; i >= 0; i-- {
			pv := p.arr[i].w
			
			lc := p.LeftChild(i)
			if size > lc {
				lcv := p.arr[lc].w
				if lcv < pv {
					p.swap(lc, i)
				}
			}

			rc := p.RightChild(i)
			if size > rc {
				rcv := p.arr[rc].w
				if rcv < pv {
					p.swap(rc, i)
				}
			}
		} 
	}
}