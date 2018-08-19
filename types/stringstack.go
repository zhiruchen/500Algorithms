package types

// StringStack string stack
type StringStack struct {
	items []string
}

func NewStringStack() *StringStack {
	return &StringStack{items: make([]string, 0)}
}
func (s *StringStack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *StringStack) Pop() string {
	if s.Empty() {
		return ""
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *StringStack) Items() []string {
	return s.items
}

func (s *StringStack) Empty() bool {
	return len(s.items) == 0
}
