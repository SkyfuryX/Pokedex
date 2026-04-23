package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello","world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c  := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %v, got length: %v", len(c.expected), len(actual))
		}
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if  word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}