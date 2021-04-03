package ta_lab6

import (
	"fmt"
	"testing"
)

func TestUnbalancedTree(t *testing.T) {
	tree := NewUnbalancedTree(&element{1})
	fmt.Println(tree)
	tree.insert(&element{2})
	fmt.Println(tree)
	tree.insert(&element{0})
	fmt.Println(tree)
	tree.insert(&element{0})
	tree.insert(&element{3})
	tree.insert(&element{4})
	tree.insert(&element{5})
	tree.insert(&element{6})
	tree.insert(&element{7})
	tree.insert(&element{8})
	fmt.Println(tree)

	fmt.Println(tree)
	fmt.Println(tree.find(&element{1}))
}
