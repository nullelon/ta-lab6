package ta_lab6

import "fmt"

type UnbalancedTree struct {
	Left  *UnbalancedTree
	Value Comparable
	Right *UnbalancedTree
}

func NewUnbalancedTree(initElement Comparable) *UnbalancedTree {
	return &UnbalancedTree{nil, initElement, nil}
}

func (t *UnbalancedTree) delete(v Comparable) {
	panic("implement me")
}

func (t *UnbalancedTree) find(v Comparable) {
	panic("implement me")
}

func (t *UnbalancedTree) insert(v Comparable) *UnbalancedTree {
	if t == nil {
		return &UnbalancedTree{nil, v, nil}
	}

	if v.Less(t.Value) {
		t.Left = t.Left.insert(v)
	} else {
		t.Right = t.Right.insert(v)
	}
	return t
}

func (t *UnbalancedTree) String() string {
	if t == nil {
		return "()"
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
