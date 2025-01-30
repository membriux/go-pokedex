package main

import (
	"testing"
)

type cleanInputTestCase struct {
	input string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []cleanInputTestCase{
		{
			input: "    hello   ",
			expected: []string{"hello"},
		},
		{
			input: "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   HellO World    ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("%s does not match %s", word, expectedWord)
				t.Fail()
			}
		}
	}


}
