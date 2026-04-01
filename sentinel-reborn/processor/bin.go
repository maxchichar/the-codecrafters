// ummulkusum musa
package main

import (
	"fmt"
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

func main() {
	fmt.Println(BinToDec("101010 (bin) mr abraham the senior dev, and our able code legend"))
	fmt.Println(BinToDec("10 (bin)"))
}
