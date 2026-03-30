// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import(
	"strings"
)

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