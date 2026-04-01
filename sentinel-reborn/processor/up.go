// Agi Ruth

package main

import (
	"fmt"
	"strings"
)

func convertUp(input string) string {
	words := strings.Fields(input)
	result := []string{
		
	}

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		} else {
			result = append(result, words[i])
		}
	}

	return strings.Join(result, " ")
}

func main() {
	// input := "Ready, set, go (up) !"
	// output := convertUp(input)
	// fmt.Println(output)

	fmt.Println(convertUp("Ready, set, go (up) !"))
}