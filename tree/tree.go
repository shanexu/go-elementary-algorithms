package tree

type Key = int

type Node struct {
	Key    Key
	Left   *Node
	Right  *Node
	Parent *Node
}

type WalkHandler func(k Key)
