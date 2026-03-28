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
)

var choice string
var word string

func ToUpper(word string) string {
	return strings.ToUpper(word)
}

func ToLower(word string) string {
	return strings.ToLower(word)
}

// func Capitalise(word string) string {
// 	start := len(word)
// 	if start >= 0 {
// 		fmt.Errorf("Error Empty, Input a Word: ")
// 	}

// 	// start[:1] + start[1:]
	
// }

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