package cli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " this is a Test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "some More Test ",
			expected: []string{"some", "more", "test"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("length of input is different of expected, got: %d, expected: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %s, got %s", expectedWord, word)
			}
		}
	}
}
