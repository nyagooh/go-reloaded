package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Convert_Binary function converts binary numbers to decimal format based on specified patterns.
func Convert_Binary(str string) string {
	words := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\((BIN|B|I|N|b|i|n|Bin|BIn|BiN|bIN)\)`)
	if pattern1.MatchString(str) {
		fmt.Println("invalid pattern found")
		return " "
	} else {
		for i, ch := range words {
			if ch == "(bin)" && i > 0 {
				number, err := strconv.ParseInt(words[i-1], 2, 64)
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
