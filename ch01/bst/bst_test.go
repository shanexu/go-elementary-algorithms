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

	a = nil
	PreOrderWalk(r, walkHandler)
	assert.EqualValues(t, []int{4, 3, 1, 2, 8, 7, 16, 10, 9, 14}, a)

	a = nil
	InOrderWalk(r, walkHandler)
	assert.EqualValues(t, []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, a)

	a = nil
	PostOrderWalk(r, walkHandler)
	assert.EqualValues(t, []int{2, 1, 3, 7, 9, 14, 10, 16, 8, 4}, a)
}
