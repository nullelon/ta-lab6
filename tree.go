package ta_lab6

import (
	"fmt"
)

type Comparable interface {
	Less(a interface{}) bool
}

type ITree interface {
	fmt.Stringer
	insert(v Comparable) *UnbalancedTree
	delete(v Comparable)
	find(v Comparable)
}
