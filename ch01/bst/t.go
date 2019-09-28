package bst

import (
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

type SimpleNode struct {
	key    Key
	left   Node
	right  Node
	parent Node
}

func (s *SimpleNode) Key() Key {
	return s.key
}

func (s *SimpleNode) SetKey(key Key) {
	s.key = key
}

func (s *SimpleNode) Left() Node {
	return s.left
}

func (s *SimpleNode) SetLeft(l Node) {
	s.left = l
}

func (s *SimpleNode) Right() Node {
	return s.right
}

func (s *SimpleNode) SetRight(r Node) {
	s.right = r
}

func (s *SimpleNode) Parent() Node {
	return s.parent
}

func (s *SimpleNode) SetParent(p Node) {
	s.parent = p
}

func CreateLeaf(k Key) Node {
	return &SimpleNode{
		key:    k,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}
