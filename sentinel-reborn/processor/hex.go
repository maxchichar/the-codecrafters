// EDWNI EJMEBI
package main

import (
	"strconv"
	"strings"
)

func HexToDecimal(bin string) string {
	s := strings.Fields(bin)
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
	}
	return strings.Join(s, " ")
}
