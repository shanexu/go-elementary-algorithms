package bst

import (
	. "github.com/shanexu/go-elementary-algorithms/ch01/tree"
)

func Insert(t *Node, k int) *Node {
	root := t
	x := CreateLeaf(k)
	var parent *Node
	for t != nil {
		parent = t
		if k < t.Key {
			t = t.Left
		} else if k == t.Key {
			return root
		} else {
			t = t.Right
		}
	}
	x.Parent = parent
	if parent == nil {
		return x
	}
	if k < parent.Key {
		parent.Left = x
	} else {
		parent.Right = x
	}
	return root
}

func CreateLeaf(k int) *Node {
	return &Node{
		Key:    k,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}
