package main

import (
	"fmt"
	"strconv"
)
var option string
var userInput string

func Hex(hexStr string) (int64, error) {
		n, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			return 0, fmt.Errorf("Invalid Hexadecimal Number")
		}
		return n, nil

}

func Bin(binStr string) (int64, error) {
		n, err := strconv.ParseInt(binStr, 2, 64)
		if err != nil {
			return 0, fmt.Errorf("Invalid Binary Number")
		}
		return n, nil
}

func DecToHex(Input string) (string, error) {
	n, err := strconv.Atoi(Input)
	if err != nil {
		return "", fmt.Errorf("Invalid Decimal Number")
	}
	return strconv.FormatInt(int64(n), 16), nil
}

func DecToBin(Input string) (string, error) {
	n, err := strconv.Atoi(Input)
	if err != nil {
		return "", fmt.Errorf("Invalid Decimal Number")
	}
	return strconv.FormatInt(int64(n), 2), nil
}

func main()  {
	fmt.Println("Welcome To The Base Converter")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("Choose conversion by the number")
		fmt.Println("1. Hexadecimal to decimal")
		fmt.Println("2. Binary to decimal")
		fmt.Println("3. Decimal to hexadecimal and binary")
		fmt.Println("4. Quit")
		
		fmt.Println()
		
		fmt.Print("Option: ")
		fmt.Scanln(&option)
		fmt.Println()
	
		switch option {
		case "1":
			fmt.Print("Input Hexadecimal: ")
			fmt.Scanln(&userInput)
			fmt.Println()
			
			fmt.Printf("> Convert %v hex\n", userInput)
			fmt.Print("Decimal: ")
			fmt.Println(Hex(userInput))
			continue
		case "2":
			fmt.Print("Input Binary: ")
			fmt.Scanln(&userInput)
			fmt.Println()

			fmt.Printf("> Convert %v bin\n", userInput)
			fmt.Print("Decimal: ")
			fmt.Println(Bin(userInput))
			continue
		case "3":
			fmt.Print("Input Decimal: ")
			fmt.Scanln(&userInput)
			fmt.Println()

			fmt.Printf("> Convert %v dec\n", userInput)
			fmt.Print("Binary: ")
			fmt.Println(DecToBin(userInput))
			fmt.Print("Hexadecimal: ")
			fmt.Println(DecToHex(userInput))
			
		case "4", "Quit", "quit":
			fmt.Println("See You Again")
			return
		default:
			fmt.Println("Error: Input invalid")
		}
		
	}

}