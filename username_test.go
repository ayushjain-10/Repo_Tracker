package main

import (
	"testing"
	"unicode"
	"strings"
)

func isValidUsername(username string) bool {
	// Username is between 1 and 39 characters
	if len(username) < 1 || len(username) > 39 {
		return false
	}
	// Username may only contain alphanumeric characters or hyphens
	for _, ch := range username {
		if !(ch == '-' || unicode.IsLetter(ch) || unicode.IsNumber(ch)) {
			return false
		}
	}
	// Username cannot have multiple consecutive hyphens
	if strings.Contains(username, "--") {
		return false
	}
	// Username cannot begin or end with a hyphen
	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return false
	}
	return true
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{"Test 1", "valid-username", true},
		{"Test 2", "not--valid", false},
		{"Test 3", "-invalid", false},
		{"Test 4", "invalid-", false},
		{"Test 5", "inval!d", false},
		{"Test 6", strings.Repeat("a", 40), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidUsername(tt.username); got != tt.want {
				t.Errorf("isValidUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
