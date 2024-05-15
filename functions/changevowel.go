package functions

import (
	"regexp"
	"strings"
)

// ChangeVowel function adds 'n' after 'a' if followed by a vowel in a string.
func ChangeVowel(str string) string {
	text := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\b[aA]\b`)
	pattern2 := regexp.MustCompile(`[aeiouhAEIOUH]`)
	for i := range str {
		if i+1 < len(text) && i > 0 && pattern1.MatchString(text[i]) && pattern2.MatchString(text[i+1]) {
			text[i] += "n"
		}
	}
	str = strings.Join(text, " ")
	return strings.TrimSpace(str)
}
