package btree

import (
	"fmt"
)

// datatype used to store keys
type key int

// maps a node to depth for printing
type keyDepth struct {
	k     key
	depth int
}

// A node represents a node in a Btree
type node struct {
	keys     []key
	children []*node
	deg      int  // minimum degree
	keyc     int  // number of keys stored currently
	leaf     bool // is this node a leaf
}

// newNode is a node constructor for a node of degree deg
// leaf indicates whether the node is a leaf
func newNode(deg int, leaf bool) *node {
	keys := make([]key, 2*deg-1) // len(keys) = 0, cap(keys) = 2d-1
	kids := make([]*node, 2*deg) // len(kids) = cap(kids) = 2d all nil
	n := node{keys, kids, deg, 0, leaf}
	return &n
}

// search n for the given tree. If key is not found
// and n is not a leaf node, search children. If key is
// not found and n is a leaf, return the node, false
func (n *node) search(k key) (*node, bool) {
	// find the index of the first key >k
	i := 0
	for ; i < len(n.keys) && n.keys[i] < k; i++ {
	}

	// key found
	if n.keys[i] == k {
		return n, true
	}

	// key not found, no more children to search
	if n.leaf {
		return n, false
	}

	// search the appropriate child
	return n.children[i].search(k)
}

// split node at given child key
func (n *node) splitChild(i int, child *node) {
	// node to be sibling of split child
	sib := newNode(child.deg, child.leaf)
	sib.keyc = n.deg - 1

	// copy last t-1 keys of child to sib
	for t, j := n.deg, 0; j < t-1; j++ {
		sib.keys[j] = child.keys[j+t]
	}

	// copy last t children of child to sib
	if child.leaf == false {
		for j, t := 0, n.deg; j < t; j++ {
			sib.children[j] = child.children[j+t]
		}
	}

	// reduce number of keys in child
	child.keyc = n.deg - 1

	// make room for new child and key in n
	for ci, ki := n.keyc, n.keyc-1; ci >= i+1 && ki >= i; ci, ki = ci+1, ki+1 {
		n.children[ci+1] = n.children[ci]
		n.keys[ki+1] = n.keys[ki]
	}
	// the new child is sib, the new key is the
	// middle key from child. insert both into n
	// and increment n.keyc
	n.children[i+1] = sib
	n.keys[i] = child.keys[n.deg-1]
	n.keyc++
}

// insertNonFull
func (n *node) insertNonFull(k key) {
	// initialize i as index of last element
	i := n.keyc - 1

	// if n is a leaf, find index for k and shift
	// all greater keys up by 1
	if n.leaf {
		for ; i >= 0 && n.keys[i] > k; i-- {
			n.keys[i+1] = n.keys[i]
		}
		// insert k
		n.keys[i+1] = k
		n.keyc++
		return
	}

	// n is not a leaf
	// find index of child where k will go
	for ; i >= 0 && n.keys[i] > k; i-- {
	}
	// if the child is full split it
	child := n.children[i+1]
	if child.full() {
		n.splitChild(i+1, child)
		// find which of the children from split
		// to insert into
		if n.keys[i+1] < k {
			child = n.children[i+2]
		}
	}
	child.insertNonFull(k)
}

// full
func (n *node) full() bool {
	return (n.keyc == 2*n.deg-1)
}

// traverse
func (n *node) traverse(d int, c chan<- *keyDepth) {
	// if n is a leaf, write all keys to the channel
	if n.leaf {
		fmt.Printf("leaf node: %v\n", n.keys[:n.keyc])
		for _, k := range n.keys[:n.keyc] {
			c <- &keyDepth{key(k), d}
		}
		return
	}
	// otherwise recursively traverse children, writing
	// the separating keys in-between
	// remember to close the channel
	fmt.Printf("non-leaf node: %v\n", n.keys[:n.keyc])
	var i int
	for ; i < n.keyc; i++ {
		n.children[i].traverse(d+1, c)
		c <- &keyDepth{key(n.keys[i]), d}
	}
	n.children[i].traverse(d+1, c)
	if d == 0 {
		close(c)
	}
}

// implement Stringer for node
func (n *node) String() string {
	return fmt.Sprintf("node: { degree: %d, leaf: %t }", n.deg, n.leaf)
}
