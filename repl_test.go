package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  pikachu  ",
			expected: []string{"pikachu"},
		},
		{
			input:    "Charmander Bulbasaur SQUIRTLE",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "RaYquAZa groUDon              KYOGRE",
			expected: []string{"rayquaza", "groudon", "kyogre"},
		},
	} // cases

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Expected Length: %d, Actual Length: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s, got %s", expectedWord, word)
			} // if
		} // for i
	} // for c
} // TestCleanInput
