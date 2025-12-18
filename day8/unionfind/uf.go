package uf

import "fmt"

type Node struct {
	inited bool // Has MakeSet been called on this item?
	//  If !inited all other fields invalid
	parent int // Parent of node, points to self if root
	rank   int // Height of tree, only valid if root
}

type UnionFind struct {
	data []Node
}

// Debug Printout of Data
func (uf *UnionFind) PrintData() {
	for i, ent := range uf.data {
		fmt.Println(i, ": ", ent)
	}
}

func (uf *UnionFind) MakeSet(x int) {
	if x >= len(uf.data) {
		temp := make([]Node, 2*x)
		copy(temp, uf.data)
		uf.data = temp
	}

	if uf.data[x].inited {
		return
	}

	uf.data[x] = Node{true, x, 0}
}

// Pre:
//
//	MakeSet(x) has been called
func (uf *UnionFind) Find(x int) int {
	curr := uf.data[x]

	if !curr.inited {
		return -1
	}

	if uf.data[curr.parent].parent != curr.parent {
		curr = uf.data[curr.parent]
	}

	return curr.parent
}

// Pre:
//
//	x, y < len(uf.data)
//	MakeSet(x) and MakeSet(y) have been called
func (uf *UnionFind) Union(x, y int) {
	lf := uf.Find(x)
	rf := uf.Find(y)

	if uf.data[lf].rank < uf.data[rf].rank {
		uf.data[lf].parent = rf
	} else if uf.data[lf].rank > uf.data[rf].rank {
		uf.data[rf].parent = lf
	} else {
		uf.data[rf].parent = lf
		uf.data[lf].rank += 1
	}
}
