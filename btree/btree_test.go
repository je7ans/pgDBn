package btree

import (
	"fmt"
)

func ExampleNode() {
	leaf := newNode(3, true)
	nleaf := newNode(4, false)
	fmt.Println(leaf)
	fmt.Println(nleaf)
	// Output:
	// node: { degree: 3, leaf: true }
	// node: { degree: 4, leaf: false }
}

func ExampleBtree() {
	bt := NewBtree(2)
	for i, _ := range []int{6, 2, 5, 9, 0, 1, 7, 3, 8, 4} {
		bt.Insert(key(i))
	}
	fmt.Println(bt)
	// Output:
	// btree: { degree: 5, root: { <nil> } }
}
