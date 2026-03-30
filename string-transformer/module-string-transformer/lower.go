// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import(
	"strings"
)

func ToLower(text string) string {
	if strings.HasPrefix(text, "lower") {
		if len(text) > 0 {
			text = strings.TrimPrefix(text, "lower") // removes lower from the output
			text = strings.ToLower(text) //converts any text to lower case
		}
	}
	return text
}