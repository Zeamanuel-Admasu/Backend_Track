package main

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "abeba",
			expected: true,
		},
		{
			input:    "broorb",
			expected: true,
		},
		{
			input:    "abebe",
			expected: false,
		},
		{
			input:    "beso",
			expected: false,
		},
	}
	for _, test := range tests {
		result := check(test.input)
		if result != test.expected {
			t.Errorf("for input %q, expected id %t but got %t", test.input, test.expected, result)
		}
	}

}
