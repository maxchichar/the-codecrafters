//janai egeonu

package main

import (
	"fmt"
	"os"
	"sentinel-reborn/processor"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("ERROR: INVALID ARGUMENTS!!")
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Fprintln(os.Stderr, "Input and output cannot be the same files")
		os.Exit(1)
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("sorry.. Error Reading input file, ERROR =", err)
		return
	}

	result := processor.Complier(string(data))

	err = os.WriteFile(outputFile, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("sorry.. Error writing output file, ERROR =", err)
	}

}
