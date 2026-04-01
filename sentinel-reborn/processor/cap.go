package processor

import (
	"strings"
)

func Capitalize(s string) string {
	w := strings.Fields(s)
	for i := 0; i < len(w); i++ {
		if w[i] == "(cap)" {
			w[i-1] = strings.ToUpper(w[i-1][:1]) + strings.ToLower(w[i-1][1:])
			w = append(w[:i], w[i+1:]...)

		}

	}
	return strings.Join(w, " ")
}
