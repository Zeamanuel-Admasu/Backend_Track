package main

import "testing"

func TestWordAmountCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{input: "go is a beautiful language!",
			expected: map[string]int{"go": 1, "is": 1, "a": 1, "beautiful": 1, "language": 1},
		},
	}
	for _, test := range tests {
		result := wordAmountCount(test.input)

		if len(result) != len(test.expected) {
			t.Errorf("for input: %q, expected %v but got %v", test.input, test.expected, result)
			continue
		}
		for key, expectedValue := range test.expected {
			if result[key] != expectedValue {
				t.Errorf("For input %q, word %q expected %d but got %d", test.input, key, expectedValue, result[key])
			}
		}

	}
}
