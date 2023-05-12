package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Given a valid phone number with special formatting (spaces, parentheses,
// hyphens), returns a normalized phone number with all the special formatting
// removed and just the 10 digits. Returns an error if the number of digits in
// the input is not exactly 10.
func normalizePhone(raw string) (string, error) {
	var normalized strings.Builder

	for _, r := range raw {
		if unicode.IsDigit(r) {
			normalized.WriteRune(r)
		}
	}

	if normalized.Len() != 10 {
		return "", fmt.Errorf("Invalid input: %s", raw)
	}

	return normalized.String(), nil
}
