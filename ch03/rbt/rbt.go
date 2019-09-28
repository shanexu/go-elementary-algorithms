package rbt

import (
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

func LeftRotate(t, x Node) Node {
	p := x.Parent()
	y := x.Right()
	a := x.Left()
	b := y.Left()
	c := y.Right()
	Replace(x, y)
	SetChildren(x, a, b)
	SetChildren(y, x, c)
	if p == nil {
		t = y
	}
	return t
}

func RightRotate(t, y Node) Node {
	p := y.Parent()
	x := y.Left()
	a := x.Left()
	b := x.Right()
	c := y.Right()
	Replace(y, x)
	SetChildren(y, b, c)
	SetChildren(x, a, y)
	if p == nil {
		t = x
	}
	return t
}

func Replace(x, y Node) {
	if x.Parent() == nil {
		if y != nil {
			y.SetParent(nil)
		}
	} else if x.Parent().Left() == x {
		SetLeft(x.Parent(), y)
	} else {
		SetRight(x.Parent(), y)
	}
	x.SetParent(nil)
}

func SetChildren(x, l, r Node) {
	SetLeft(x, l)
	SetRight(x, r)
}

func SetLeft(x, y Node) {
	x.SetLeft(y)
	if y != nil {
		y.SetParent(x)
	}
}

func SetRight(x, y Node) {
	x.SetRight(y)
	if y != nil {
		y.SetParent(x)
	}
}

