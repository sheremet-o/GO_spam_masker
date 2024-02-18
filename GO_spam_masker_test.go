package main

import "testing"

func TestMasker(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "Here's my spammy page: http://hehefouls.netHAHAHA see you.", expected: "Here's my spammy page: http://******************* see you."},
	}

	for _, testCase := range testCases {
		result := masker(testCase.input)
		if result != testCase.expected {
			t.Errorf("Должно быть %s", testCase.expected)
		}
	}
}