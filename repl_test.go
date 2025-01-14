package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "	hello world	",
			expected: []string{"hello", "world"},
		},
		{
			input:    "	heLLo there 	",
			expected: []string{"hello", "there"},
		},
		{
			input:    "           ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Failed")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Failed")
			}
		}
	}

}
