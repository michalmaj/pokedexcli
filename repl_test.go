package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "trim_and_collapse_spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "lowercase_all",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:     "tabs_and_newlines",
			input:    "\n\t  Mew\tTwo \t ",
			expected: []string{"mew", "two"},
		},
		{
			name:     "empty_string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only_spaces",
			input:    "    ",
			expected: []string{},
		},
		{
			name:     "no_split_on_non_whitespace",
			input:    "hello-world",
			expected: []string{"hello-world"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				t.Fatalf("wrong length: got %d, want %d; actual=%v expected=%v",
					len(actual), len(c.expected), actual, c.expected)
			}

			for i := range c.expected {
				if actual[i] != c.expected[i] {
					t.Errorf("word %d mismatch: got %q, want %q (input=%q)",
						i, actual[i], c.expected[i], c.input)
				}
			}
		})
	}
}