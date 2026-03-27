package main

import (
	"fmt"
	"strconv"
)
var option int
var userInput string

func Hex(hexStr string) (int64, error) {
	for{
		n, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			fmt.Print("Invalid input. Input a valid hex: ")
			fmt.Scanln(&userInput)
			continue
		}
		return n, nil
	}

}

func Bin(binStr string) (int64, error) {
	for {
		n, err := strconv.ParseInt(binStr, 2, 64)
		if err != nil {
			fmt.Print("Invalid input. Input a number: ")
			fmt.Scanln(&userInput)
			continue
		}
		return n, nil
	}
}
/*
func DecToHex(decHStr int64, base int) string {
	n, err := strconv.FormatInt(decHStr, 16)
	if err != nil {
		return 0, err)
	}
	return n, nil
}
*/

func main()  {
	fmt.Println("Welcome To The Base Converter")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("Select an operation by the number")
		op1 := 1
		fmt.Println(op1, ". Convert hexadecimal to decimal")
		op2 := 2
		fmt.Println(op2, ". Convert binary to decimal")
		op3 := 3
		fmt.Println(op3, ". Convert decimal to hexadecimal and binary")
		op4 := 4
		fmt.Println(op4, ". Quit")
		
		fmt.Println()
		
		fmt.Print("Option: ")
		fmt.Scanln(&option)
		fmt.Println()
	
		switch option {
		case 1:
			fmt.Print("Input the hexadecimal you want to convert: ")
			fmt.Scanln(&userInput)
			fmt.Println()
			
			fmt.Printf("> Convert %v hex\n", userInput)
			fmt.Print("Decimal: ")
			fmt.Println(Hex(userInput))
			continue
		case 2:
			fmt.Println("Input the binary you want to convert: ")
			fmt.Scanln(&userInput)
			fmt.Println()

			fmt.Printf("> Convert %v bin\n", userInput)
			fmt.Print("Decimal: ")
			fmt.Println(Bin(userInput))
			continue
		case 3:
			fmt.Println("Input the decimal you want to convert: ")
			fmt.Scanln(&userInput)
			fmt.Println()

			fmt.Printf("> Convert %v dec\n", userInput)
			fmt.Print("Binary: ")
			
		case 4:
			fmt.Println("See You Again")
			return
		}
	}

}