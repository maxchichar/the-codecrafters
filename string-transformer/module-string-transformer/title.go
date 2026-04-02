// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package modulestringtransformer

import(
	"strings"
)

func Title(text string) string {
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