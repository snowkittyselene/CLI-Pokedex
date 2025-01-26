package main

import (
	"testing"
)

func TestCleanIput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "Testing testing testing",
			expected: []string{"testing", "testing", "testing"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Failed test %s: %d does not match expected length %d", c.input, len(actual), len(c.expected))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Failed: %s does not match expected word %s at position %d", word, expectedWord, i)
				t.Fail()
			}
		}

	}
}
