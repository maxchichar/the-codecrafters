// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"strings"
)

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