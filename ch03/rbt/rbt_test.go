package rbt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shanexu/go-elementary-algorithms/ch01/bst"
	. "github.com/shanexu/go-elementary-algorithms/tree"
)

func NewTree(ks []Key) Node {
	var r Node
	for _, k := range ks {
		r = bst.Insert(r, k)
	}
	return r
}

func TestRBT(t *testing.T) {
	ks := []Key{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	r := NewTree(ks)

	r = LeftRotate(r, r)
	var a []Key
	InOrderWalk(r, func(k Key) {
		a = append(a, k)
	})
	assert.Equal(t, []Key{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, a)

	r = RightRotate(r, r)
	a = nil
	InOrderWalk(r, func(k Key) {
		a = append(a, k)
	})
	assert.Equal(t, []Key{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, a)
}
