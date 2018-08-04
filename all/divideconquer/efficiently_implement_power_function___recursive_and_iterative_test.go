package divideconquer

import (
	"testing"
)

func TestEfficientlyImplementPowerFunctionRecursiveAndIterative(t *testing.T) {
	cases := []struct {
		x     int
		n     uint
		power int64
	}{
		{
			x:     2,
			n:     3,
			power: 8,
		},
		{
			x:     2,
			n:     4,
			power: 16,
		},
		{
			x:     -2,
			n:     3,
			power: -8,
		},
		{
			x:     -2,
			n:     10,
			power: 1024,
		},
	}

	for _, c := range cases {
		power := EfficientlyImplementPowerFunctionRecursiveAndIterative(c.x, c.n)
		if power != c.power {
			t.Errorf("Expect: %d, get: %d\n", c.power, power)
		}
	}
}
