package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfGivenExpressionIsBalancedExpressionOrNot(t *testing.T) {
	ast := assert.New(t)
	cases := []struct {
		exp string
		b   bool
	}{
		{
			exp: "([{}])",
			b:   true,
		},
		{
			exp: "([])",
			b:   true,
		},
		{
			exp: "([{()}])",
			b:   true,
		},
		{
			exp: "([{{])",
			b:   false,
		},
		{
			exp: "((]}))",
			b:   false,
		},
	}

	for _, c := range cases {
		result := CheckIfGivenExpressionIsBalancedExpressionOrNot(c.exp)
		ast.Equal(c.b, result, fmt.Sprintf("result not equal %t, get: %t", c.b, result))
	}
}
