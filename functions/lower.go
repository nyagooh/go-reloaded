package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// IsLower function processes a string to ensure lowercase modifications based on defined patterns.
func IsLower(str string) string {
	text := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\((LOW|L|O|W|l|o|w|LOw|LoW|loW|lOW)\)`)
	if pattern1.MatchString(str) {
		fmt.Println("invalid pattern found")
		return " "
	} else {
		for i, ch := range text {
			if ch == "(low)" && i > 0 {
				text[i-1] = strings.ToLower(text[i-1])
				text = append(text[:i], text[i+1:]...)
			}
			if ch == "(low," && i > 0 {
				words := strings.Trim(text[i+1], ")")
				number, err := strconv.Atoi(words)
				if err != nil {
					log.Fatalf("atoi has an error")
				}
				if number > len(text) {
					log.Fatalf("lower number is greater than the length of the string")
				}
				for k := 1; k <= number; k++ {
					if i-k < 0 {
						log.Fatalf("number out of scope")
					}
					text[i-k] = strings.ToLower(text[i-k])
				}
				text = append(text[:i], text[i+2:]...)
			}
		}
	}
	results := strings.Join(text, " ")
	return results
}

