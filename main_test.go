package main

import (
	"testing"
)

func TestIsNameInvalid(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
	}{
		{"valid_name.txt", false},
		{"valid-name.txt", false},
		{"valid name.txt", false},
		{"invalid<name.txt", true},
		{"invalid>name.txt", true},
		{"invalid:name.txt", true},
		{"invalid\"name.txt", true},
		{"invalid|name.txt", true},
		{"invalid*name.txt", true},
		{"invalid?name.txt", true},
		{"invalid/name.txt", true},
		{".hiddenfile.txt", true},
	}

	for _, testCase := range testCases {
		result := isNameInvalid(testCase.name)
		if result != testCase.expected {
			t.Errorf("isNameInvalid('%s') returned %v, expected %v", testCase.name, result, testCase.expected)
		}
	}
}

func TestReplaceInvalidCharacters(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{"valid_name.txt", "valid_name.txt"},
		{"invalid<name.txt", "invalid_name.txt"},
		{"invalid>name.txt", "invalid_name.txt"},
		{"invalid:name.txt", "invalid_name.txt"},
		{"invalid\"name.txt", "invalid_name.txt"},
		{"invalid|name.txt", "invalid_name.txt"},
		{"invalid*name.txt", "invalid_name.txt"},
		{"invalid?name.txt", "invalid_name.txt"},
		{"invalid/name.txt", "invalid_name.txt"},
		{"file#number.txt", "fileNRnumber.txt"},
	}

	for _, testCase := range testCases {
		result := replaceInvalidCharacters(testCase.name)
		if result != testCase.expected {
			t.Errorf("replaceInvalidCharacters('%s') returned '%s', expected '%s'", testCase.name, result, testCase.expected)
		}
	}
}
