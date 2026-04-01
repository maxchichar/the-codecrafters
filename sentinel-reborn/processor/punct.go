// Name: Agene Okoh

package processor

import (
	"strings"
)

func FixPunctuation(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	
		s = strings.ReplaceAll(s, " ,", ",")
		s = strings.ReplaceAll(s, " .", ".")
		s = strings.ReplaceAll(s, " !", "!")
		s = strings.ReplaceAll(s, " ?", "?")
		s = strings.ReplaceAll(s, " :", ":")
		s = strings.ReplaceAll(s, " ;", ";")
		
	
	return s

}

func HasPunctuation(s string) bool {
	return strings.ContainsAny(s,",.?:!")

}
