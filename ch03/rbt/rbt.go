package rbt

import (
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

func LeftRotate(t *Node, x *Node) *Node {
	y := x.Right
	b := y.Left
	p := x.Parent
	x.Right = b
	b.Parent = x
	x.Parent = y
	y.Left = x
	y.Parent = p
	if p == nil {
		return y
	}
	if p.Left == x {
		p.Left = y
	} else {
		p.Right = y
	}
	return t
}

func RightRotate(t *Node, x *Node) *Node {
	panic("// TODO") // TODO
}
