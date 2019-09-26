package bst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/shanexu/go-elementary-algorithms/ch01/bst"
	. "github.com/shanexu/go-elementary-algorithms/ch01/tree"
)

func TestInsert(t *testing.T) {
	var r *Node

	ks := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}

	for _, k := range ks {
		r = Insert(r, k)
	}

	var a []int

	var walkHandler WalkHandler = func(k int) {
		a = append(a, k)
	}

	preOrder := []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}
	a = nil
	PreOrderWalk(r, walkHandler)
	assert.Equal(t, preOrder, a)

	inOrder := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	a = nil
	InOrderWalk(r, walkHandler)
	assert.Equal(t, inOrder, a)

	postOrder := []int{2, 1, 3, 7, 9, 14, 10, 16, 8, 4}
	a = nil
	PostOrderWalk(r, walkHandler)
	assert.Equal(t, postOrder, a)

	min := Min(r)
	assert.Equal(t, 1, min.Key)

	max := Max(r)
	assert.Equal(t, 16, max.Key)

	n := Search(r, 2)
	n = Succ(n)
	assert.Equal(t, 3, n.Key)

	n = Search(r, 8)
	n = Succ(n)
	assert.Equal(t, 9, n.Key)

	n = Search(r, 16)
	n = Pred(n)
	assert.Equal(t, 14, n.Key)

	n = Search(r, 9)
	n = Pred(n)
	assert.Equal(t, 8, n.Key)

	x := ReConstructFromPreOrderAndInOrder(preOrder, inOrder)
	a = nil
	PostOrderWalk(x, walkHandler)
	assert.Equal(t, postOrder, a)
}
