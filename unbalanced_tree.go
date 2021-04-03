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
	panic("implement me")
}

func (t UnbalancedTree) find(v Comparable) interface{} {
	if t.Value == nil {
		return nil
	}

	if v.Equals(t.Value) {
		return t.Value
	}

	if v.Less(t.Value) {
		return t.Left.find(v)
	} else {
		return t.Right.find(v)
	}
}

func (t *UnbalancedTree) insert(v Comparable) *UnbalancedTree {
	if t.Value == nil {
		return NewUnbalancedTree(v)
	}

	if t.Value.Equals(v) {
		return t
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
