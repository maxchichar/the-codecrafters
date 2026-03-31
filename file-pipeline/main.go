// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceTodo(text string) string {
	return strings.ReplaceAll(text, "TODO:", "✦ ACTION:")
}

func trimSpaces(text string) string {
	return strings.TrimSpace(text)
}

func replaceClassified(text string) string {
	if strings.HasPrefix(text, "CLASSIFIED:") {
		return "[REDACTED]:" + text[len("CLASSIFIED:"):]
	}
	return text
}

func allCapsToTitle(text string) string {
	if text == strings.ToUpper(text) {
		return strings.Title(strings.ToLower(text))
	}
	return text
}

func lowerToUpper(text string) string {
	if text == strings.ToLower(text) {
		return strings.ToUpper(text)
	}
	return text
}

// func removeBlanks(text string) string {
// 	trimmed := strings.TrimSpace(text)
// 	if trimmed == "" || isAllDashes(trimmed) {
// 		return ""
// 	}
// 	return text
// }


func main()  {
	//Error & egde case handling
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	if inputPath == outputPath {
		fmt.Fprintln(os.Stderr, "✗ Input and output cannot be the same file")
		os.Exit(1)
	}

	file, err := os.Open(inputPath) // for read access
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ File not found: %s\n", file)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	const maxCapacity = 3000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
}