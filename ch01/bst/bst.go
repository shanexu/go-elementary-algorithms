package bst

import (
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

func Insert(t Node, k Key) Node {
	root := t
	x := CreateLeaf(k)
	var parent Node
	for t != nil {
		parent = t
		if k < t.Key() {
			t = t.Left()
		} else if k == t.Key() {
			return root
		} else {
			t = t.Right()
		}
	}
	x.SetParent(parent)

	if parent == nil {
		return x
	}
	if k < parent.Key() {
		parent.SetLeft(x)
	} else {
		parent.SetRight(x)
	}
	return root
}

func Search(t Node, x Key) Node {
	for t != nil && x != t.Key() {
		if x < t.Key() {
			t = t.Left()
		} else {
			t = t.Right()
		}
	}
	return t
}

func Min(t Node)Node {
	for ; t.Left() != nil; t = t.Left() {
	}
	return t
}

func Max(t Node) Node {
	for ; t.Right() != nil; t = t.Right() {
	}
	return t
}

func Succ(x Node) Node {
	if x.Right() != nil {
		return Min(x.Right())
	}
	p := x.Parent()
	for p != nil && x == p.Right() {
		x = p
		p = x.Parent()
	}
	return p
}

func Pred(x Node) Node {
	if x.Left() != nil {
		return Max(x.Left())
	}
	p := x.Parent()
	for p != nil && x == p.Left() {
		x = p
		p = x.Parent()
	}
	return p
}

func reConstructFromPreOrderAndInOrder(preOrder []Key, inOrder []Key, parent Node) Node {
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
	x.SetParent(parent)
	x.SetLeft(reConstructFromPreOrderAndInOrder(preOrder[1:i+1], inOrder[0:i], x))
	x.SetRight(reConstructFromPreOrderAndInOrder(preOrder[i+1:], inOrder[i+1:], x))
	return x
}

func ReConstructFromPreOrderAndInOrder(preOrder []Key, inOrder []Key) Node {
	return reConstructFromPreOrderAndInOrder(preOrder, inOrder, nil)
}

func Delete(t, x Node) Node {
	r := t
	p := x.Parent()
	x0 := x

	if x.Left() == nil {
		x = x.Right()
	} else if x.Right() == nil {
		x = x.Left()
	} else {
		y := Min(x.Right())
		x.SetKey(y.Key())
		if y.Parent() != x {
			y.Parent().SetLeft(y.Right())
		} else {
			x.SetRight(y.Right())
		}
		if y.Right() != nil {
			y.Right().SetParent(y.Parent())
		}
		return r
	}
	if x != nil {
		x.SetParent(p)
	}
	if p == nil {
		r = x
	} else {
		if p.Left() == x0 {
			p.SetLeft(x)
		} else {
			p.SetRight(x)
		}
	}
	return r
}

func DeleteL(t, x Node) Node {
	r := t
	p := x.Parent()
	x0 := x

	if x.Left() == nil {
		x = x.Right()
	} else if x.Right() == nil {
		x = x.Left()
	} else {
		y := Max(x.Left())
		x.SetKey(y.Key())
		if y.Parent() != x {
			y.Parent().SetRight(y.Left())
		} else {
			x.SetLeft(y.Left())
		}
		if y.Left() != nil {
			y.Left().SetParent(y.Parent())
		}
		return r
	}
	if x != nil {
		x.SetParent(p)
	}
	if p == nil {
		r = x
	} else {
		if p.Left() == x0 {
			p.SetLeft(x)
		} else {
			p.SetRight(x)
		}
	}
	return r
}

