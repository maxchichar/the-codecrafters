// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import(
	"strings"
)

func ToUpper(text string) string {
	if strings.HasPrefix(text, "upper") {
		text = strings.TrimPrefix(text, "upper")
		text = strings.ToUpper(text)
	}
	return text
}