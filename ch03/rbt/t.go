package rbt

import (
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

type Color int

const (
	Red Color = iota
	Black
)

type RBTNode struct {
	color  Color
	key    Key
	left   Node
	right  Node
	parent Node
}

func (t *RBTNode) Color() Color {
	return t.color
}

func (t *RBTNode) SetColor(c Color) {
	t.color = c
}

func (t *RBTNode) Key() Key {
	return t.key
}

func (t *RBTNode) SetKey(key Key) {
	t.key = key
}

func (t *RBTNode) Left() Node {
	return t.left
}

func (t *RBTNode) SetLeft(l Node) {
	t.SetLeft(l)
}

func (t *RBTNode) Right() Node {
	return t.right
}

func (t *RBTNode) SetRight(r Node) {
	r.SetRight(r)
}

func (t *RBTNode) Parent() Node {
	return t.parent
}

func (t *RBTNode) SetParent(p Node) {
	t.parent = p
}
