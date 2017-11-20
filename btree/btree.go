// Package btree provides a B-Tree data structure
package btree

import (
	"bytes"
	"strconv"
	"strings"
)

// A Btree data structure
type Btree struct {
	root *node // root node
	deg  int   // minimum degree
}

// Btree constructor
func NewBtree(deg int) *Btree {
	t := Btree{nil, deg}
	return &t
}

// search
func (t *Btree) search(k key) (*node, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.search(k)
}

// insert
func (t *Btree) Insert(k key) {
	// if root is nil create node with single key
	// and set as root
	if t.root == nil {
		t.root = newNode(t.deg, true)
		t.root.keys[0] = k
		t.root.keyc = 1
		return
	}

	// if root is full, create new node with root as its child
	// split root and set the new node to be root. finally insert
	// k into the appropriate child
	// otherwise call insertNonFull with root
	if t.root.full() {
		s := newNode(t.root.deg, false)
		s.children[0] = t.root
		s.splitChild(0, t.root)
		var i int
		if s.keys[0] < k {
			i++
		}
		s.children[i].insertNonFull(k)
		t.root = s
	} else {
		t.root.insertNonFull(k)
	}
}

// implement Stringer for Btree
func (t *Btree) String() string {
	kds := make(chan *keyDepth)
	go t.root.traverse(0, kds)
	var pre string
	var buffer bytes.Buffer
	for kd := range kds {
		k, d := kd.k, kd.depth
		pre = strings.Repeat("--", d)
		buffer.WriteString(pre)
		buffer.WriteString(strconv.Itoa(int(k)))
		buffer.WriteString("\n")
	}
	return buffer.String()
}
