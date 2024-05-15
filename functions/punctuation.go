package functions

import (
	"regexp"
	"strings"
)

// Punctuation removes extra spaces around punctuation marks in a string using regex.
func Punctuation(str string) string {
	pattern := regexp.MustCompile(`\s*([,.?!:;]+)\s*`)
	pattern2 := pattern.ReplaceAllString(str, "$1 ")
	return strings.TrimSpace(pattern2)
}
