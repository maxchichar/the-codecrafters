// EDWNI EJMEBI
// ummulkusum musa

package processor

import (
	"strconv"
	"strings"
)

func Base(text string) string {

	s := strings.Fields(text)
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" && i > 0 {
			x, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				continue
			}
			s[i-1] = strconv.FormatInt(x, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		}

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
