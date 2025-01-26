package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello    world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "       Hello,      this is    a   test. You're    funny!",
			expected: []string{"Hello,", "this", "is", "a", "test.", "You're", "funny!"},
		},
		{
			input: "  I     am  a     big     fan     of Jax!",
			expected: []string{"I", "am", "a", "big", "fan", "of", "Jax!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Length does not match, please try again")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s does not match %s, please try again", word, expectedWord)
			}
		}
	}
}