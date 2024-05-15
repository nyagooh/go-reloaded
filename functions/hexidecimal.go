package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// HexiDecimal function converts hexadecimal numbers to decimal format based on specified patterns.
func HexiDecimal(str string) string {
	words := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\((HEX|H|X|E|h|x|e|Hex|heX|hEx|HEx|HEx)\)`)
	if pattern1.MatchString(str) {
		fmt.Println("invalid pattern found")
		return " "
	} else {
		for i, ch := range words {
			if ch == "(hex)" && i > 0 {
				number, err := strconv.ParseInt(words[i-1], 16, 64)
				if err != nil {
					log.Fatalf(err.Error())
				}
				words[i-1] = strconv.FormatInt(number, 10)
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	return strings.Join(words, " ")
}
