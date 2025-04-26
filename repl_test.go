package main

import "testing"

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "zeRgliNg         hYdrAliSk         zEalOt",
			expected: []string{"zergling", "hydralisk", "zealot"},
		},
	}

	for _, test := range testCases {
		result := cleanInput(test.input)

		if len(result) != len(test.expected) {
			t.Fatalf("received %d words, expected %d", len(result), len(test.expected))
		}

		for i := range result {
			word := result[i]
			expectedWord := test.expected[i]
			if word != expectedWord {
				t.Fatalf("receieved %s word, expected %s", word, expectedWord)
			}

		}
	}
}
