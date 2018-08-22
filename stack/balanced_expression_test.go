package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalancedExpression(t *testing.T) {
	cases := []struct {
		exp string
		b   bool
	}{
		{
			exp: "{[()]}",
			b:   true,
		},
		{
			exp: "{[([)]}",
			b:   false,
		},
		{
			exp: "{[([])]",
			b:   false,
		},
	}

	for _, c := range cases {
		b := IsBalancedExpression(c.exp)
		assert.Equal(t, c.b, b, "exp is not balance")
	}
}
