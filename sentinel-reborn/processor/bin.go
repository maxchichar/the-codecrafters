// ummulkusum musa
package processor

import (
	"strconv"
	"strings"
)

func BinToDec(bin string) string {

	s := strings.Fields(bin)
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" && i > 0 {
			conv, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				continue
			}
			s[i-1] = strconv.FormatInt(conv, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return strings.Join(s, " ")
}
