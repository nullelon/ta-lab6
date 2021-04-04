package ta_lab6

import (
	"fmt"
	"testing"
)

func TestUnbalancedTree(t *testing.T) {
	tree := NewUnbalancedTree(&element{50})
	fmt.Println(tree)
	tree.insert(&element{25})
	fmt.Println(tree)
	tree.insert(&element{75})
	fmt.Println(tree)
	tree.insert(&element{30})
	tree.insert(&element{31})
	tree.insert(&element{33})
	tree.insert(&element{21})
	tree.insert(&element{20})
	tree.insert(&element{22})
	tree.insert(&element{23})

	fmt.Println(tree)
	tree.delete(&element{21})

	fmt.Println(tree)
}

func TestRBTree(t *testing.T) {
	tree := NewRedBlackTree(&element{50})
	fmt.Println(tree)

	tree.insert(&element{5})
	fmt.Println(tree)

	tree.insert(&element{10})
	fmt.Println(tree)

	tree.insert(&element{15})
	fmt.Println(tree)

	tree.insert(&element{7})
	fmt.Println(tree)

	tree.insert(&element{4})
	fmt.Println(tree)

	tree.insert(&element{2})
	fmt.Println(tree)

	tree.insert(&element{8})
	fmt.Println(tree)

	tree.insert(&element{7})
	fmt.Println(tree)

	tree.insert(&element{222})
	fmt.Println(tree)

	tree.insert(&element{0})
	fmt.Println(tree)
}
