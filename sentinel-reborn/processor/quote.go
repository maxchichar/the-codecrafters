// Chibueze Maxwell

package processor

import(
	"strings"
)

func fixQuote(text string) string {
	word := strings.Split(text, "'")
	for i, value := range word {
		word[i] = strings.TrimSpace(value)
	}

	return strings.Join(word, "'")
}
