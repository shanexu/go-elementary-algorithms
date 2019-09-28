package tree

type Key = int

type Node interface {
	Key() Key
	SetKey(key Key)
	Left() Node
	SetLeft(l Node)
	Right() Node
	SetRight(r Node)
	Parent() Node
	SetParent(p Node)
}

type WalkHandler func(k Key)

func PreOrderWalk(t Node, f WalkHandler) {
	if t == nil {
		return
	}
	f(t.Key())
	PreOrderWalk(t.Left(), f)
	PreOrderWalk(t.Right(), f)
}

func InOrderWalk(t Node, f WalkHandler) {
	if t == nil {
		return
	}
	InOrderWalk(t.Left(), f)
	f(t.Key())
	InOrderWalk(t.Right(), f)
}

func PostOrderWalk(t Node, f WalkHandler) {
	if t == nil {
		return
	}
	PostOrderWalk(t.Left(), f)
	PostOrderWalk(t.Right(), f)
	f(t.Key())
}
