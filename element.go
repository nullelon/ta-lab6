package ta_lab6

type element struct {
	value int
}

func (e *element) Less(a interface{}) bool {
	return e.value < a.(element).value
}
