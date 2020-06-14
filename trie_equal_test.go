package trie_test

import (
	"testing"

	"github.com/reflechant/trie"
)

/* func TestTwoNilEqual(t *testing.T) {
	var n1, n2 *Node
	if !Equal(n1, n2) {
		t.FailNow()
	}
} */

/* func TestOnlyOneNil(t *testing.T) {
	var n1, n2 *Node = new(Node), nil
	if Equal(n1, n2) {
		t.FailNow()
	}
} */
func TestEmptyEqual(t *testing.T) {
	var n1, n2 trie.Node
	if !trie.Equal(n1, n2) {
		t.FailNow()
	}
}

func TestDiffValsNonEqual(t *testing.T) {
	n1 := trie.Node{KeyPart: 'й'}
	n2 := trie.Node{KeyPart: 'f'}
	if trie.Equal(n1, n2) {
		t.FailNow()
	}
}

func TestDiffLenNonEqual(t *testing.T) {
	// n1 := trie.Node{Val: 'f'}
	// n2 := trie.Node{Val: 'f', leafs: []Node{Node{}}}
	// if trie.Equal(n1, n2) {
	// 	t.FailNow()
	// }
}

func TestDiffOrderEqual(t *testing.T) {
	// if !trie.Equal(n1, n2) {
	// 	t.FailNow()
	// }
}
