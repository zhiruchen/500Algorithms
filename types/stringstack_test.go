package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringStack(t *testing.T) {
	ast := assert.New(t)
	cases := []struct {
		items []string
		val   string
	}{
		{
			items: []string{"1", "2", "3"},
			val:   "3",
		},
	}

	for _, c := range cases {
		s := NewStringStack()
		for _, v := range c.items {
			s.Push(v)
		}

		ast.False(s.Empty(), "stack is empty")
		ast.Equal([]string{"1", "2", "3"}, s.Items(), "stack items nor equal")
		ast.Equal(s.Pop(), c.val, "pop value not equal")
		ast.Equal([]string{"1", "2"}, s.Items(), "stack items nor equal")
	}
}
