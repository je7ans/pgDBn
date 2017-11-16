// Package btree provides a B-Tree data structure
package btree

// datatype used to store keys
type key int

// A node represents a node in a Btree
type node struct {
	keys     []key
	children []*node
	deg      int  // minimum degree
	leaf     bool // is this node a leaf
}

// newNode is a node constructor for a node of degree deg
// leaf indicates whether the node is a leaf
func newNode(deg int, leaf bool) *node {
	keys := make([]key, 0, 2*d-1) // len(keys) = 0, cap(keys) = 2d-1
	kids := make([]*node, 2*d)    // len(kids) = cap(kids) = 2d all nil
	n := node{keys, kids, deg, leaf}
	return &n
}

// search n for the given tree. If key is not found
// and n is not a leaf node, search children. If key is
// not found and n is a leaf, return nil
func (n *node) search(k key) *node {
	// find the index of the first key >k
	i := 0
	for ; i < len(n.keys) && keys[i] < k; i++ {
	}

	// key found
	if keys[i] == k {
		return n
	}

	// key not found, no more children to search
	if n.leaf {
		return nil
	}

	// search the appropriate child
	return n.children[i].search(k)
}

// implement Stringer for node
func (n *node) String() string {}

// A Btree data structure
type Btree struct {
	root *node // root node
	deg  int   // minimum degree
}

// TODO
func NewBtree() *Btree {}

// search
func (t *Btree) search(k key) *Btree {
	if t.root == nil {
		return nil
	}
	return t.root.search(key)
}

// implement Stringer for Btree
func (n *node) String() string {}
