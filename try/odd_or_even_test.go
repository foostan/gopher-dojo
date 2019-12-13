package try

import (
	"testing"
)

func TestIsEven(t *testing.T) {

	t.Helper()

	cases := []struct {name string; input int; expected bool}{
		{name: "+odd", input: 5, expected: false},
		{name: "+even", input: 6, expected: true},
		{name: "-odd", input: -5, expected: false},
		{name: "-even", input: -6, expected: true},
		{name: "zero", input: 0, expected: false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := IsEven(c.input); c.expected != actual {
				t.Fatalf("want IsEven(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}

}
