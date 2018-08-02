package stack

import (
	"fmt"

	stack "github.com/emirpasic/gods/stacks/arraystack"
)

// DuplicateParenthesis check expression contains any duplicate parenthesis or not
func DuplicateParenthesis(exp string) {
	s := stack.New()

	for _, x := range []rune(exp) {
		if x != ')' {
			s.Push(x)
		} else {
			vv, _ := s.Peek()
			v := vv.(rune)
			if v == '(' {
				fmt.Println("find duplicate parenthesis")
				return
			}

			for v != '(' {
				s.Pop()
				vv, _ = s.Peek()
				v = vv.(rune)
			}

			s.Pop()
		}
	}

	fmt.Println(exp, " doest not have duplicate parenthesis")
}
