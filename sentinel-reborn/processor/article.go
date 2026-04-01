// Blessing Anebi
// Emmanuel Inogwu
package processor

import (
	"strings"
)

func FixArticles(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {

		if word[i] == "a" || word[i] == "A" && strings.ContainsRune("aeiouhAEIOUH", rune(word[i+1][0])) {
			word[i] = "An"
		}
	}

	return strings.Join(word, " ")
}
