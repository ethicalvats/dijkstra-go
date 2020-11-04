package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sort"
)

type TNode struct {
	v int
	w float64
}

type MinPath struct {
	n int
	d float64
}

type Tree struct {
	queue PriorityQueue
	nodes map[int][]TNode
	d map[int]float64
}

func (t *Tree) Start(size int) {
	t.Dijkstra(size)
}

func (t *Tree) Dijkstra(size int) {

	//STEP: initialize(G, s)
	d := make(map[int]float64, size)

	for i := 1; i <= size; i++ {
		d[i] = math.Inf(1)
	}

	d[1] = 0 // d value of source 0

	pq := PriorityQueue{}

	// priority Queue
	for j := 0; j < size; j++ {
		pq.arr = append(pq.arr, TNode{j+1, d[j+1]})
	}

	t.queue = pq
	t.d = d

	// empty set S <- nil
	visit := make(map[int]bool)

	//STEP: while Q != nil
	for len(pq.arr) > 0 {

		//STEP: U <- EXTRACT_MIN(Q)
		u := pq.ExtractMin()

		//STEP: S <- S u {U}
		visit[u.v] = true

		//STEP: foreach verter adj to s
		for _, n := range t.nodes[u.v] {
			
			//STEP: Relax(u, v, w)
			dv := t.Relax(u.v, n.v, n.w)

			t.d[n.v] = dv
			t.queue.Update(n.v, dv)
		}

	}

	output(t.d)
}

func output(d map[int]float64){
	var k []int
	var O []string
	
	for v := range d {
		if v != 1 { //ignoring the source vertex d value
			k = append(k, v)
		}
	}

	// sorting keys as the order is not deterministic
	sort.Ints(k)

	for _,i := range k {
		O = append(O, strconv.Itoa(int(d[i])))
	}
	
	fmt.Println(strings.Join(O, " "))
}


// if d[v] > d[u] + w(u,v) -> d[v] = d[u] + w(u,v)
func (t *Tree) Relax(u int, v int, w float64) float64  {

	dv := t.d[v]
	du := t.d[u]

	if dv > du + w {
		dv = du + w
	}

	return dv
}

func main(){

	// fmt.Println("Main")

	n1 := TNode{2, 5}
	n2 := TNode{3, 2}
	n3 := TNode{4, 1}
	n4 := TNode{4, 6}
	n5 := TNode{5, 5}


	tree := Tree{}
	tree.nodes = make(map[int][]TNode)

	tree.nodes[1] = []TNode{n1, n2, n4}
	tree.nodes[3] = []TNode{n3, n5}

	tree.Start(5)
}