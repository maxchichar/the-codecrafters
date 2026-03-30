// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import (
	"regexp"
	"strings"
)

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