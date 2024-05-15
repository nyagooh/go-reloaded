package functions

import (
	"regexp"
	"strings"
)

// Apostrophe function removes extra spaces around apostrophes in a string.
func Apostrophe(str string) string {
	pattern1 := regexp.MustCompile(`'\s*(.*?)\s'`)
	pattern2 := pattern1.ReplaceAllString(str, "'$1' ")
	return strings.TrimSpace(string(pattern2))
}
