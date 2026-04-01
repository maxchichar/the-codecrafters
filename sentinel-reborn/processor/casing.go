// Faith Ochanya Ejembi
// Agi Ruth

package processor

import (
	"strconv"
	"strings"
)

func Casing(text string) string {

	words := strings.Fields(text)

	for i := len(words) - 1; i >= 0; i-- {
		if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}

		if strings.HasPrefix(words[i], "(low,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}
		if strings.HasPrefix(words[i], "(cap,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(string(words[i-j][0])) + strings.ToLower(words[i-j][1:])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}

		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)

		}

		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

	}
	return strings.Join(words, " ")
}
