// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"regexp"
)

var choice string
var word string

func ToCapital(text string) string {
	s := strings.ToLower(text) //converts the text to lowercase
	strings.Fields(text)
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
	if strings.HasPrefix(text, "lower") {
		if len(text) > 0 {
			text = strings.TrimPrefix(text, "lower") // removes lower from the output
			text = strings.ToLower(text) //converts any text to lower case
		}
	}
	return text
}

func ToUpper(text string) string {
	if strings.HasPrefix(text, "upper") {
		text = strings.TrimPrefix(text, "upper")
		text = strings.ToUpper(text)
	}
	return text
}

func Title(text string) string {
	words := strings.Fields(text)
	smallwords := " a an the and but or for nor on at to by in of up as is it "
	
	
	for i, w := range words {
		if strings.Contains(smallwords, " "+w+" ") {
			words[i] = w
		}else{
			words[i] = strings.Title(w)
		}
	}
	
	if len(words) > 0 && strings.ToLower(words[0]) == "title" {
		strings.TrimPrefix(words[0], "title")
		words = words[1:]
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
		strings.TrimPrefix(text, "reverse")
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

func main()  {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
  	fmt.Println("──────────────────────────────────────")
	  fmt.Println()

	fmt.Print("S.S.T Loading...")
  	time.Sleep(3 * time.Second)
	fmt.Println()
	  
	for {
		fmt.Println()
		fmt.Println("Menu Loading...")
		time.Sleep(2 * time.Second)
		fmt.Println()
		
		fmt.Println("Sentinel String Transformer Menu")
		fmt.Println("────────────────────────────────")
		fmt.Println("⚠️ Select option using numbers ❗")
		set1 := 1 
		fmt.Println(set1, "To uppercase")
		set2 :=	2
		fmt.Println(set2, "To lowercase")
		set3 := 3
		fmt.Println(set3, "Capitalize")
		set4 := 4
		fmt.Println(set4, )
		set5 := 5
		fmt.Println(set5)
		set6 := 6
		fmt.Println(set6)
		set7 := 7
		fmt.Println(set7, "Exit")

		fmt.Println()
		fmt.Print("Enter option: ")
		fmt.Scanln(&choice)
		
		option, err := strconv.Atoi(choice)
		if err == nil {
			continue
		}
		switch choice {
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			fmt.Println("Please select the option by number: ")
			fmt.Scanln(&choice)
			continue
		case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
			fmt.Println("Please select the option by number: ")
			fmt.Scanln(&choice)
			continue
		}
		
		if option == 1 {
			fmt.Println("Welcome to Sentinel Uppercase Transformer")
			fmt.Println()

			fmt.Println("Instructions Loading...")
			time.Sleep(1 * time.Second)
			fmt.Println()

			fmt.Print("User Input instruction: ")
			fmt.Println("upper <text>")
			
			
			// upper <text>
		}
		if option == 7{
			return
		}
		
	}


  
}