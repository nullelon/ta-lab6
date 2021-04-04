package ta_lab6

//for all insert cases look at en.wikipedia.org/wiki/Red–black_tree

import (
	"fmt"
)

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

func NewRedBlackTreeFromNode(initNode *Node) *RedBlackTree {
	return &RedBlackTree{Root: initNode}
}

func NewRedNode(value Comparable, ancestor *Node) *Node {
	return &Node{
		Value:    value,
		Ancestor: ancestor,
		Color:    RED,
	}
}

func (t *RedBlackTree) delete(v Comparable) {
	panic("implement me")
}

func (t *RedBlackTree) find(v Comparable) interface{} {
	return t.Root.find(v)
}

func (n *Node) find(v Comparable) interface{} {
	if n.Value == nil {
		return nil
	}

	if v.Equals(n.Value) {
		return n.Value
	}

	if v.Less(n.Value) && n.Left != nil {
		return n.Left.find(v)
	} else if n.Right != nil {
		return n.Right.find(v)
	} else {
		return nil
	}
}

func (t *RedBlackTree) insert(v Comparable) ITree {
	n := t.Root.insert(v)
	n.repairTree(t)
	return t
}

func (n *Node) insert(v Comparable) *Node {
	if n.Value == nil {
		panic("insert value is nil")
	}

	if v.Equals(n.Value) {
		return n
	}

	if v.Less(n.Value) {
		if n.Left == nil {
			n.Left = NewRedNode(v, n)
			return n.Left
		} else {
			return n.Left.insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = NewRedNode(v, n)
			return n.Right
		} else {
			return n.Right.insert(v)
		}
	}
}

func (n *Node) repairTree(t *RedBlackTree) {
	if n.Ancestor == nil {
		t.Root = n
		return
	}
	if n.Ancestor.Color == BLACK {
		return
	} else { //Ancestor Color is red -> red violation
		if n.GrandFather() == nil {
			n.Ancestor.Color = BLACK
			return
		}
		if n.Uncle() == nil || n.Uncle().Color != RED { //Insert case 4 & 5

			//~refactor
			if (n == n.Ancestor.Right && n.Ancestor == n.GrandFather().Left) || (n == n.Ancestor.Left && n.Ancestor == n.GrandFather().Right) {

				isAncestorLeft := n.Ancestor == n.GrandFather().Left
				n = n.rotateDir(n.Ancestor, isAncestorLeft)
			}

			n.Ancestor.Color = BLACK
			n.GrandFather().Color = RED

			n.RotateDirRoot(t) // G may be the root
		} else { //Insert case 1
			n.Ancestor.Color = BLACK
			n.Uncle().Color = BLACK
			n.GrandFather().Color = RED

			n.GrandFather().repairTree(t)
		}
	}
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
	s += fmt.Sprint([]string{"R", "B"}[n.Color], n.Value)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}

func (n *Node) Uncle() *Node {
	isAncestorLeft := n.Ancestor == n.GrandFather().Left
	if isAncestorLeft {
		return n.GrandFather().Right
	} else {
		return n.GrandFather().Left
	}
}

func (n *Node) GrandFather() *Node {
	return n.Ancestor.Ancestor
}

func (n *Node) RotateDirRoot(t *RedBlackTree) {
	n = n.Ancestor
	if n.GrandFather() == nil {
		t.Root = n
	} else {

		isAncestorLeft := n.Ancestor == n.GrandFather().Left

		if isAncestorLeft {
			n.GrandFather().Left = n
		} else {
			n.GrandFather().Right = n
		}
	}

	isNLeft := n == n.Ancestor.Left

	if isNLeft {
		n.Ancestor.Left = n.Right
		n.Right = n.Ancestor

		g := n.GrandFather()
		n.Ancestor.Ancestor = n
		n.Ancestor = g
	} else {
		n.Ancestor.Right = n.Left
		n.Left = n.Ancestor

		g := n.GrandFather()
		n.Ancestor.Ancestor = n
		n.Ancestor = g
	}

}

func (n *Node) rotateDir(ancestor *Node, isAncestorLeft bool) *Node {
	if isAncestorLeft {
		n.GrandFather().Left = n
		ancestor.Right = n.Left
		n.Left = ancestor

		n.Ancestor = n.GrandFather()
		ancestor.Ancestor = n
	} else {
		n.GrandFather().Right = n
		ancestor.Left = n.Right
		n.Right = ancestor

		n.Ancestor = n.GrandFather()
		ancestor.Ancestor = n
	}

	return ancestor

}
