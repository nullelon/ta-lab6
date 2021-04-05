package ta_lab6

import (
	"fmt"
	"github.com/emirpasic/gods/utils"
)

type color int

const (
	RED = iota
	BLACK
)

// RedBlackTree holds elements of the RED-BLACK tree
type RedBlackTree struct {
	Root       *Node
	Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key    interface{}
	Value  interface{}
	Color  color
	Left   *Node
	Right  *Node
	Parent *Node
}

// NewWithIntComparator instantiates a RED-BLACK tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *RedBlackTree {
	return &RedBlackTree{Comparator: utils.IntComparator}
}

// insert inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTree) insert(key interface{}) {
	value := key
	var insertedNode *Node
	if tree.Root == nil {
		// Assert key is of comparator's type for initial tree
		tree.Comparator(key, key)
		tree.Root = &Node{Key: key, Value: value, Color: RED}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.Comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node{Key: key, Value: value, Color: RED}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node{Key: key, Value: value, Color: RED}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
}

// find searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTree) find(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// delete remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *RedBlackTree) delete(key interface{}) {
	var child *Node
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pRED := node.Left.maximumNode()
		node.Key = pRED.Key
		node.Value = pRED.Value
		node = pRED
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.Color == BLACK {
			node.Color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.Color = BLACK
		}
	}
}

func (tree *RedBlackTree) lookup(key interface{}) *Node {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *RedBlackTree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *RedBlackTree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *RedBlackTree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *RedBlackTree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.Color = BLACK
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlackTree) insertCase2(node *Node) {
	if nodeColor(node.Parent) == BLACK {
		return
	}
	tree.insertCase3(node)
}

func (tree *RedBlackTree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == RED {
		node.Parent.Color = BLACK
		uncle.Color = BLACK
		node.grandparent().Color = RED
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlackTree) insertCase4(node *Node) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *RedBlackTree) insertCase5(node *Node) {
	node.Parent.Color = BLACK
	grandparent := node.grandparent()
	grandparent.Color = RED
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *RedBlackTree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RedBlackTree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == RED {
		node.Parent.Color = RED
		sibling.Color = BLACK
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RedBlackTree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == BLACK &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RedBlackTree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == RED &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		node.Parent.Color = BLACK
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RedBlackTree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == RED &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		sibling.Left.Color = BLACK
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Right) == RED &&
		nodeColor(sibling.Left) == BLACK {
		sibling.Color = RED
		sibling.Right.Color = BLACK
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RedBlackTree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.Color = nodeColor(node.Parent)
	node.Parent.Color = BLACK
	if node == node.Parent.Left && nodeColor(sibling.Right) == RED {
		sibling.Right.Color = BLACK
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == RED {
		sibling.Left.Color = BLACK
		tree.rotateRight(node.Parent)
	}
}

func nodeColor(node *Node) color {
	if node == nil {
		return BLACK
	}
	return node.Color
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
