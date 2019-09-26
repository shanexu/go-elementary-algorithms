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

type WalkHandler func(k int)

func PreOrderWalk(t *Node, f WalkHandler) {
	if t == nil {
		return
	}
	f(t.Key)
	PreOrderWalk(t.Left, f)
	PreOrderWalk(t.Right, f)
}

func InOrderWalk(t *Node, f WalkHandler) {
	if t == nil {
		return
	}
	InOrderWalk(t.Left, f)
	f(t.Key)
	InOrderWalk(t.Right, f)
}

func PostOrderWalk(t *Node, f WalkHandler) {
	if t == nil {
		return
	}
	PostOrderWalk(t.Left, f)
	PostOrderWalk(t.Right, f)
	f(t.Key)
}

func Search(t *Node, x int) *Node {
	for t != nil && x != t.Key {
		if x < t.Key {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	return t
}

func Min(t *Node) *Node {
	for ; t.Left != nil; t = t.Left {
	}
	return t
}

func Max(t *Node) *Node {
	for ; t.Right != nil; t = t.Right {
	}
	return t
}

func Succ(x *Node) *Node {
	if x.Right != nil {
		return Min(x.Right)
	}
	p := x.Parent
	for p != nil && x == p.Right {
		x = p
		p = x.Parent
	}
	return p
}

func Pred(x *Node) *Node {
	if x.Left != nil {
		return Max(x.Left)
	}
	p := x.Parent
	for p != nil && x == p.Left {
		x = p
		p = x.Parent
	}
	return p
}

func ReConstructFromPreOrderAndInOrder(preOrder []int, inOrder []int) *Node {
	if len(preOrder) == 0 {
		return nil
	}
	rootKey := preOrder[0]
	i := 0
	for ; i < len(inOrder); i++ {
		if inOrder[i] == rootKey {
			break
		}
	}
	x := CreateLeaf(rootKey)
	x.Left = ReConstructFromPreOrderAndInOrder(preOrder[1:i+1], inOrder[0:i])
	x.Right = ReConstructFromPreOrderAndInOrder(preOrder[i+1:], inOrder[i+1:])

	return x
}

func CreateLeaf(k int) *Node {
	return &Node{
		Key:    k,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}
