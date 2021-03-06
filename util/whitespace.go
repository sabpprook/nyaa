package util

import (
	"strings"
)

// return true if r is a whitespace rune
func IsWhitespace(r rune) bool {
	return r == '\n' || r == '\t' || r == '\r' || r == ' '
}

// trim whitespaces from a string
func TrimWhitespaces(s string) string {
	s = strings.TrimLeftFunc(s, IsWhitespace)
	s = strings.TrimRightFunc(s, IsWhitespace)
	return s
}
