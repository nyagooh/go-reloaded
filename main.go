package main

import (
	"fmt"
	"log"
	"os"
	m "reloaded/functions"

)


func main(){
	//retrieve command line arguments from the terminal
	args := os.Args[1:]
	if len(args) == 0{
		log.Fatalf("no arguments provided")
		return
	}  
	// checks if files retrieved from the argument are two
	if len(args) == 2 {
		input := args[0]
		output := args[1]
		text,err := os.ReadFile(string(input))
		if err != nil {
			fmt.Println(err)
		}
		new := EditedText(string(text))
	 os.WriteFile(output,[]byte(new),0o644)
	} else {
		log.Fatalf("not enough arguments provided")
	}
	
}
// a function that calls my other program function
func EditedText (str string)string{
	var output ,result string
	 output = m.IsUpper(str)
	result = m.IsLower(output)
	result = m.Capitalize(result)
	result = m.Convert_Binary(result)
	result = m.HexiDecimal(result)
	result = m.ChangeVowel(result)
	result = m.Punctuation(result)
	result = m.Apostrophe(result)
	// if result !=
	return result

}