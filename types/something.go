package types

// SomethingStack is linear data structure
type SomethingStack struct {
	items []Something
}

func NewSomethingStack() *SomethingStack {
	return &SomethingStack{items: make([]Something, 0)}
}
func (q *SomethingStack) Push(item Something) {
	q.items = append(q.items, item)
}
func (q *SomethingStack) Pop() Something {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
