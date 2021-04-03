package ta_lab6

import "strconv"

type element struct {
	value int
}

func (e *element) Less(a interface{}) bool {
	return e.value < a.(*element).value
}

func (e *element) Equals(a interface{}) bool {
	return e.value == a.(*element).value
}

func (e element) String() string {
	return strconv.Itoa(e.value)
}
