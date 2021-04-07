package ta_lab6

import (
	"fmt"
	"math/rand"
	"testing"
)

func createUnbalancedTree(n int, isRand bool) (*UnbalancedTree, []int) {
	var tree = NewUnbalancedTree(&element{0})
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		if isRand {
			v := rand.Int()
			d[i] = v
			tree.insert(&element{v})
		} else {
			d[i] = i
			tree.insert(&element{i})
		}
	}
	return tree, d
}

func createBalancedTree(n int, isRand bool) (*RedBlackTree, []int) {
	var tree = NewWithIntComparator()
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		if isRand {
			v := rand.Int()
			d[i] = v
			tree.insert(v)
		} else {
			d[i] = i
			tree.insert(i)
		}
	}
	return tree, d
}

//441888
//229169
func BenchmarkCreateUnbalancedTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createBalancedTree(1000, false)
	}
}

func BenchmarkUnbalancedTreeInsertionGrad(b *testing.B) {
	var tree, _ = createUnbalancedTree(1000, false)
	for i := 0; i < b.N; i++ {
		tree.insert(&element{rand.Int()})
	}
}

func BenchmarkUnbalancedTreeInsertionRand(b *testing.B) {
	var tree, _ = createUnbalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.insert(&element{rand.Int()})
	}
}

func BenchmarkBalancedTreeInsertionGrad(b *testing.B) {
	var tree, _ = createBalancedTree(1000, false)
	for i := 0; i < b.N; i++ {
		tree.insert(rand.Int())
	}
}

func BenchmarkBalancedTreeInsertionRand(b *testing.B) {
	var tree, _ = createBalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.insert(rand.Int())
	}
}

func BenchmarkUnbalancedTreeDeletionRand(b *testing.B) {
	var tree, data = createUnbalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.delete(&element{data[rand.Intn(len(data)-1)]})
	}
}

func BenchmarkBalancedTreeDeletionGrad(b *testing.B) {
	var tree, data = createBalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.delete(rand.Intn(len(data) - 1))
	}
}

func BenchmarkUnbalancedTreeSearchRand(b *testing.B) {
	var tree, data = createUnbalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.find(&element{data[rand.Intn(len(data)-1)]})
	}
}

func BenchmarkBalancedTreeSearchGrad(b *testing.B) {
	var tree, data = createBalancedTree(1000, true)
	for i := 0; i < b.N; i++ {
		tree.find(rand.Intn(len(data) - 1))
	}
}

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
	tree := NewWithIntComparator()
	fmt.Println(tree)

	tree.insert(5)
	fmt.Println(tree)

	tree.insert(10)
	fmt.Println(tree)

	tree.insert(15)
	fmt.Println(tree)

	tree.insert(7)
	fmt.Println(tree)

	tree.insert(4)
	fmt.Println(tree)

	tree.insert(2)
	fmt.Println(tree)

	tree.insert(8)
	fmt.Println(tree)

	tree.insert(7)
	fmt.Println(tree)

	tree.insert(0)
	fmt.Println(tree)

	fmt.Println("___")
	fmt.Println(tree)

	tree.delete(7)
	fmt.Println(tree)

	tree.delete(10)
	fmt.Println(tree)

	tree.delete(2)
	fmt.Println(tree)

	tree.delete(8)
	fmt.Println(tree)
}
