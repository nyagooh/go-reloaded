package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Capitalize function capitalizes words based on specified patterns.
func Capitalize(str string) string {
	words := strings.Fields(str)
	pattern1 := regexp.MustCompile(`\((CAP|C|A|P|c|a|p|Cap|cAp|caP|CAp|CAp)\)`)
	if pattern1.MatchString(str) {
		fmt.Println("invalid pattern found")
		return " "
	} else {
		for i, ch := range words {
			if ch == "(cap)" && i > 0 {
				words[i-1] = cases.Title(language.English).String(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
			if ch == "(cap," && i > 0 {
				text := strings.Trim(words[i+1], ")")
				number, err := strconv.Atoi(text)
				if err != nil {
					log.Fatalf("error in atoi")
				}
				if number > len(words) {
					log.Fatalf("number is greater than the length of the string")
				}
				for k := 1; k <= number; k++ {
					if i-k < 0 {
						log.Fatalf("number out of range")
					}
					words[i-k] = cases.Title(language.English).String(words[i-k])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	result := strings.Join(words, " ")
	return result
}
