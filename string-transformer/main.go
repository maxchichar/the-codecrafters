// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)


func ToCapital(text string) string {
	s := strings.ToLower(text) //converts the text to lowercase
	// As the name implies prefix: it checks for the prefix before execution
	if strings.HasPrefix(s, "cap") {
		s = strings.TrimPrefix(s, "cap")
		if len(s) > 0 {
			s = strings.Title(s)
		}
	}
	return s
}

func ToLower(text string) string {
	text = strings.ToLower(text)
	if strings.HasPrefix(text, "lower") {
		if len(text) > 0 {
			text = strings.TrimPrefix(text, "lower") // removes lower from the output
			text = strings.ToLower(text) //converts any text to lower case
		}
	}
	return text
}

func ToUpper(text string) string {
	text = strings.ToLower(text)
	if strings.HasPrefix(text, "upper") {
		text = strings.TrimPrefix(text, "upper")
		text = strings.ToUpper(text)
	}
	return text
}

func Title(text string) string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	smallwords := " a an the and i my but or for nor on at to by in of up as is it "
	
	if len(words) == 0 {
		return ""
	}
	
	for i, w := range words {
		if strings.Contains(smallwords, " "+w+" ") {
			words[i] = w
		}else{
			words[i] = strings.Title(w)
		}
	}
	
	if len(words) > 0 && strings.ToLower(words[0]) == "title" {
		words = words[1:]
	}

	if len(words) == 0 {
		return ""
	}

	words[0] = strings.Title(words[0])
	return strings.Join(words, " ")
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
func snakeCase(text string) string {
	text = strings.ToLower(text)
	text = nonAlphanumericRegex.ReplaceAllString(text, "")

	word := strings.Fields(text)

	if len(word) > 0 && strings.ToLower(word[0]) == "snake" {
		word = word[1:]
	}
	
	return strings.Join(word, "_")
}

func Reverse(text string) string {
	text = strings.ToLower(text)

	if strings.HasPrefix(text, "reverse") && strings.ToLower(text) == "reverse" {
		text = strings.TrimPrefix(text, "reverse")
	}

	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ReversedText(input string) string {
	words := strings.Fields(input)

	reversedWords := []string{}

	for _, word := range words {
		reversedWords = append(reversedWords, Reverse(word))
	}

	return strings.Join(reversedWords, " ")
}

func showHelp() {
	fmt.Println("Sentinel String Transformer Help")
	fmt.Println("──────────────────────────────────────")
	fmt.Println()

	helptext := `
	Usage: 
		<command> <text>
	Command:
		> upper     <text> 
		> snake     <text> 
		> lower     <text> 
		> cap 	    <text> 
		> reverse   <text> 
		> title 	<text>
		> exit
	Example:
		> upper sentinel online
		> SENTINEL ONLINE

		> exit
		> Shutting down String Transformer...

	`
	fmt.Print(helptext)

}

func main()  {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
  	fmt.Println("──────────────────────────────────────")
	fmt.Println()

	fmt.Print("S.S.T Loading...")
  	time.Sleep(3 * time.Second)
	fmt.Println()
	  

	fmt.Println()
	fmt.Println("Menu Loading...")
	time.Sleep(2 * time.Second)
	fmt.Println()
		
	fmt.Println("Sentinel String Transformer Menu")
	fmt.Println("────────────────────────────────")
	fmt.Println("Input 'help' to start")
	fmt.Println()
	
	scanner := bufio.NewScanner(os.Stdin)

	start:
	fmt.Print("> ")
	scanner.Scan() 
	// if !scanner.Scan() {
	// 	if err := scanner.Err(); err != nil{
	// 		fmt.Println("Error Reading Input: ", err)
	// 	}
	// 	return
	// }
	input := scanner.Text()
	input = strings.TrimSpace(input)

	word := strings.Fields(input)

	if len(word) == 0 {
		fmt.Println("✗ No text provided. Usage: 'help' to start")
		goto start
	}

	command := strings.ToLower(word[0])
	
	if len(word) == 1 && command == "upper" || command == "lower" || command == "cap" || command == "title" || command == "snake" || command == "reverse" {
		fmt.Printf("Error: '%s' requires text input. Type 'help' for examples.\n", command)
		goto start
	}

	if len(word) == 1 && command != "help" && command != "exit" {
		fmt.Printf("Unknown command. Type 'help' to see available commands.\n")
		goto start
	}
	
	
	switch command {
	case "help", "HELP", "Help":
		showHelp()
		goto start
	case "upper", "UPPER", "Upper":
		fmt.Println(ToUpper(input))
		goto start
	case "lower", "LOWER", "Lower":
		fmt.Println(ToLower(input))
		goto start
	case "cap", "CAP", "Command":
		fmt.Println(ToCapital(input))
		goto start
	case "title", "TITLE", "Title":
		fmt.Println(Title(input))
		goto start
	case "snake", "SNAKE", "Snake":
		fmt.Println(snakeCase(input))
		goto start
	case "reverse", "REVERSE", "Reverse":
		fmt.Println(ReversedText(input))
		goto start
	case "exit", "EXIT", "Exit":
		fmt.Println("Shutting down String Transformer...")
		time.Sleep(2 * time.Second)
		fmt.Println("\nGood bye")
		return
	}
	goto start
}
