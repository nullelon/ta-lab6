package ta_lab6

import "fmt"

type color int

const (
	RED = iota
	BLACK
)

type RedBlackTree struct {
	Root *Node
}

type Node struct {
	Left     *Node
	Value    Comparable
	Right    *Node
	Ancestor *Node
	Color    color
}

func NewRedBlackTree(initElement Comparable) *RedBlackTree {
	return &RedBlackTree{Root: &Node{
		Value: initElement, Color: BLACK,
	}}
}

func (t *RedBlackTree) delete(v Comparable) {
	panic("implement me")
}

func (t *RedBlackTree) find(v Comparable) interface{} {
	return t.Root.find(v)
}

func (t *Node) find(v Comparable) interface{} {
	if t.Value == nil {
		return nil
	}

	if v.Equals(t.Value) {
		return t.Value
	}

	if v.Less(t.Value) && t.Left != nil {
		return t.Left.find(v)
	} else if t.Right != nil {
		return t.Right.find(v)
	} else {
		return nil
	}
}

func (t *RedBlackTree) insert(v Comparable) *RedBlackTree {
	panic("implement me")
}

func (t *RedBlackTree) String() string {
	root := t.Root
	return root.String()
}

func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	if n.Value == nil {
		return ""
	}

	s := ""
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprint(n.Value)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}
