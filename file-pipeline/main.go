// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Chibueze Maxwell]
// Squad:  [Goroutines]

 // ═══════════════════════════════════════════

// SQUAD PIPELINE CONTRACT
// Squad: Goroutines
// ───────────────────────────────────────────
//
// Input line types:
//
//   1. Normal report lines
//   2. Lines in ALL CAPS
//   3. Lines in all lowercase
//   4. Lines starting with TODO:
//   5. Lines starting with CLASSIFIED:
//   6. Lines that are only dashes or blank
//   7. Lines with leading/trailing whitespace
//   8. Lines containing numbers and symbols
//
// Transformation rules (in order):
//
//   1. Trim all leading and trailing whitespace
//   2. Remove lines that are only dashes or blank
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
//
//   Header: yes — SENTINEL FIELD REPORT — PROCESSED
//   Line numbering format: 001., 002., 003. (three-digit zero-padded)
//   Summary block: yes — Lines Processed, Lines Written, Lines Removed
//
// Terminal summary fields:
//
//   ✦ Lines read    : <number>
//   ✦ Lines written : <number>
//   ✦ Lines removed : <number>
//   ✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines
// ═══════════════════════════════════════════
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

func reverseWords(text string) string {
	if strings.Contains(text, "REVERSE") {
		words := strings.Fields(text)
		for i, j := 0, len(words)-1; i < j; i, j = i + 1, j - 1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return text
}

func isAllDashes(s string) bool {
	for _, ch := range s{
		if ch != '-'{
			return false
		}
	}
	return true
}

func removeDashes(text string) string {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" || isAllDashes(trimmed) {
		return ""
	}
	return trimmed
}

func addNumber(n string) string {
	lineCounter++
	return fmt.Sprintf("%03d. %s", lineCounter, n)
}

var SelectedRules = []int{0, 1, 2, 3, 4, 5}

var lineCounter = 0

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

	inputFile, err := os.Open(inputPath) // for read access
	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ File not found: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "✗ Cannot write to output %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(inputFile)
	
	const maxCapacity = 3000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	fmt.Fprintln(writer, "SENTINEL FIELD REPORT — PROCESSED")
	linesRead, linesWritten, linesRemoved := 0, 0, 0

	RuleFunc := []func(string) string{
		replaceTodo,
		trimSpaces,
		replaceClassified,
		removeDashes,
		reverseWords,
		addNumber,
	}

	for scanner.Scan() {
		raw := scanner.Text()
		linesRead++
		line := raw

		line = trimSpaces(line) //rule 1

		for _, rid := range SelectedRules {
			line = RuleFunc[rid](line)
			if line == "" {
				linesRemoved++
				break
			}
		}
		if line == ""{
			continue
		}

		fmt.Fprintln(writer, line)
		linesWritten++
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "✗ Error reading input: %v\n", err)
		os.Exit(1)
	}

	if linesRead == 0 {
		fmt.Fprintln(os.Stderr, "Input file is empty. Nothing to process.")
	}

	writer.Flush() // Writes to output.

	// Terminal Output
	fmt.Printf("✦ Lines read: %d\n", linesRead)
	fmt.Printf("✦ Lines Written: %d\n", linesWritten)
	fmt.Printf("✦ Lines Removed: %d\n", linesRemoved)
	fmt.Printf("✦ Rules Applied: %v\n", SelectedRules)
}