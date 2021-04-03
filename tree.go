package ta_lab6

import (
	"fmt"
)

type Comparable interface {
	Less(a interface{}) bool
	Equals(a interface{}) bool
}

type ITree interface {
	fmt.Stringer
	insert(v Comparable) ITree
	delete(v Comparable)
	find(v Comparable) interface{}
}
