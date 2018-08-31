package stack

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// CheckIfGivenExpressionIsBalancedExpressionOrNot http://www.techiedelight.com/check-given-expression-balanced-expression-not/
func CheckIfGivenExpressionIsBalancedExpressionOrNot(exp string) bool {
	s := types.NewRuneStack()

	for _, c := range exp {
		if c == '[' || c == '(' || c == '{' {
			s.Push(c)
		}

		if c == ']' || c == '}' || c == ')' {
			if s.Empty() {
				return false
			}

			v := s.Pop()
			result := v == ')' && c != '('
			result = result || (v == '}' && c != '{')
			result = result || (v == ']' && c != '[')
			if result {
				return false
			}
		}
	}

	return s.Empty()
}
