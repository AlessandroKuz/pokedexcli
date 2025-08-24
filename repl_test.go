package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "ChaRmANDer",
			expected: []string{"charmander"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "Charmander   BulbaSAUR       PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Charmander\tBulbaSAUR       PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "hello\tworld", // tab between words
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\nworld", // newline between words
			expected: []string{"hello", "world"},
		},
		{
			input:    " \t hello \n world \r ", // mixed whitespace
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\t\t\tworld", // multiple tabs
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input: %q - len of actual (%d) is different from expected (%d)", c.input, len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Input: %q - words not matching at index %d: expected %s but got %s instead", c.input, i, word, expectedWord)
			}
		}
	}
}
