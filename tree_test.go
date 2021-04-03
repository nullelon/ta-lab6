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
