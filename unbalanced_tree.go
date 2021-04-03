package ta_lab6

import "fmt"

type UnbalancedTree struct {
	Left  *UnbalancedTree
	Value Comparable
	Right *UnbalancedTree
}

func NewUnbalancedTree(initElement Comparable) *UnbalancedTree {
	return &UnbalancedTree{new(UnbalancedTree), initElement, new(UnbalancedTree)}
}

func (t *UnbalancedTree) delete(v Comparable) {
	tree := t.findTree(v)
	if tree == nil {
		return
	}
	if tree.Left.isEmpty() && tree.Right.isEmpty() {
		tree.Value = nil
		return
	}
	if tree.Left.isEmpty() && !tree.Right.isEmpty() {
		tree.Value = tree.Right.Value
		return
	}
	if !tree.Left.isEmpty() && tree.Right.isEmpty() {
		tree.Value = tree.Left.Value
		return
	}

	lessParent := tree
	less := tree.Left
	fmt.Println(less.isEmpty())
	fmt.Println(less.Right.isEmpty())
	fmt.Println(less.Right.Right.isEmpty())
	for {
		if !less.isEmpty() && !less.Right.isEmpty() && less.Right.Right != nil {
			lessParent = less
			less = less.Right
		} else {
			break
		}
	}

	lessParent.Right = less.Left

	tree.Value = less.Value

}

func (t *UnbalancedTree) findTree(v Comparable) *UnbalancedTree {
	if t.Value == nil {
		return nil
	}

	if v.Equals(t.Value) {
		return t
	}

	if v.Less(t.Value) {
		return t.Left.findTree(v)
	} else {
		return t.Right.findTree(v)
	}
}

func (t *UnbalancedTree) find(v Comparable) interface{} {
	tree := t.findTree(v)
	if tree == nil {
		return nil
	}
	return tree.Value
}

func (t *UnbalancedTree) insert(v Comparable) ITree {
	if t.Value == nil {
		return NewUnbalancedTree(v)
	}

	if t.Value.Equals(v) {
		return t
	}

	if v.Less(t.Value) {
		t.Left = t.Left.insert(v).(*UnbalancedTree)
	} else {
		t.Right = t.Right.insert(v).(*UnbalancedTree)
	}
	return t
}

func (t *UnbalancedTree) isEmpty() bool {
	return t == new(UnbalancedTree)
}

func (t *UnbalancedTree) String() string {
	if t == nil {
		return "()"
	}
	if t.Value == nil {
		return ""
	}

	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
