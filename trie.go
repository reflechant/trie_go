package trie

import "container/heap"

type nodeSet []Node

func (ns nodeSet) Len() int           { return len(ns) }
func (ns nodeSet) Less(i, j int) bool { return ns[i].KeyPart < ns[j].KeyPart }
func (ns nodeSet) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns *nodeSet) Push(x interface{}) {
	node := x.(Node)
	*ns = append(*ns, node)
}

func (ns *nodeSet) Pop() interface{} {
	old := *ns
	n := len(old)
	x := old[n-1]
	*ns = old[0 : n-1]
	return x
}

// Node represents a trie node
type Node struct {
	Val     int
	KeyPart rune
	leafs   nodeSet
}

// func (n *Node) Leafs() {}
// func (n *Node) AddLeaf(r rune) {
// 	heap.Push(&n.leafs, r)
// }

// Equal returns true if trie n1 equals trie n2. Two nil tries are considered equal
func Equal(n1, n2 Node) bool {
	q1 := make([]Node, 0)
	q1 = append(q1, n1)
	q2 := make([]Node, 0)
	q2 = append(q2, n2)
	for len(q1) > 0 && len(q2) > 0 {
		n1, n2 := q1[0], q2[0]
		q1 = q1[1:]
		q2 = q2[1:]
		if n1.KeyPart != n2.KeyPart {
			return false
		}
		for len(n1.leafs) > 0 {
			q1 = append(q1, heap.Pop(&n1.leafs).(Node))
		}
		for len(n2.leafs) > 0 {
			q1 = append(q1, heap.Pop(&n2.leafs).(Node))
		}
	}
	if len(q1) != len(q2) {
		return false
	}
	return true
}
