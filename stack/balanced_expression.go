package stack

import (
	stack "github.com/emirpasic/gods/stacks/arraystack"
)

// IsBalancedExpression check if a exp is a balanced expression
func IsBalancedExpression(exp string) bool {
	s := stack.New()
	for _, x := range []rune(exp) {
		if x == '(' || x == '{' || x == '[' {
			s.Push(x)
		}

		if x == ')' || x == '}' || x == ']' {
			if s.Empty() {
				return false
			}

			vv, _ := s.Pop()
			v := vv.(rune)
			if v == '(' && x != ')' || v == '{' && x != '}' || v == '[' && x != ']' {
				return false
			}
		}
	}

	return s.Empty()
}
