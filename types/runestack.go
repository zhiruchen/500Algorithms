package types

// RuneStack string stack
type RuneStack struct {
	items []rune
}

func NewRuneStack() *RuneStack {
	return &RuneStack{items: make([]rune, 0)}
}
func (s *RuneStack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *RuneStack) Pop() rune {
	if s.Empty() {
		return 0
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *RuneStack) Items() []rune {
	return s.items
}

func (s *RuneStack) Empty() bool {
	return len(s.items) == 0
}
