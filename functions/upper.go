package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"


)

// IsUpper function processes a string to ensure uppercase modifications based on defined patterns.
func IsUpper(str string) string {
	words := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\((UP|p|u|U|P|Up|uP)\)`)
	if pattern1.MatchString(str) {
		fmt.Println("invalid pattern found")
		return " "
	} else {
		for i, ch := range words {
			if ch == "(up)" && i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
			if ch == "(up," && i > 0 {
				text := strings.Trim(words[i+1], ")")
				number, err := strconv.Atoi(text)
				if err != nil {
					fmt.Println(err)
				}
				if number > len(words) {
					log.Fatalf("upper number is greater than the length of the string")
				}
				for k := 1; k <= number; k++ {
					if i-k < 0 {
						log.Fatalf("number out of range")
					}
					words[i-k] = strings.ToUpper(words[i-k])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	result := strings.Join(words, " ")
	return result
}
