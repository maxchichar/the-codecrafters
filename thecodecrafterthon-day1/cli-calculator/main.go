package main

import (
	"fmt"
	"strconv"
)

var option string
var input1 string
var input2 string
var decision string

// Operators functions that runs the math logic

func Add(a, b int64) int64{
	return a + b
}

func Sub(a, b int64) int64 {
	return a - b
}

func Mul(a, b int64) int64 {
	return a * b
}

func Div(a, b float64) float64 {
	// error handling
	for {
		if b == 0 {
			fmt.Println("division by zero error, input another number:")
			fmt.Scanln(&b)
			continue
		}
		break
	}

	return a / b
}

func main()  {
	fmt.Println("Welcome To The Cli Calculator")
	fmt.Println()

	for {
		
		// For the first number
		fmt.Println()
		fmt.Print("Input the First number:")
		fmt.Scan(&input1)
		fmt.Println()

		num1, err := strconv.Atoi(input1)
		if err != nil {
			fmt.Println("Error: Enter a valid number")
			fmt.Println()
			continue
		}

		// For the second number
		fmt.Print("Input the Second number:")
		fmt.Scan(&input2)
		fmt.Println()

		num2, err := strconv.Atoi(input2)
		if err != nil {
			fmt.Println("Error: Enter a valid number")
			fmt.Println()
			continue
		}
	

		fmt.Println("Select an Operation by the number")

		op1 := "1. Addition"
		fmt.Println(op1)

		op2 := "2. Subtraction"
		fmt.Println(op2)

		op3 := "3. Multiplication"
		fmt.Println(op3)

		op4 := "4. Division"
		fmt.Println(op4)

		op5 := "5. Help"
		fmt.Println(op5)

		op6 := "6. Quit"
		fmt.Println(op6)
		fmt.Println()
		
		fmt.Print("Option: ")
		fmt.Scanln(&option)
		fmt.Println()

		switch option {
		case "1":
			fmt.Printf("Result = %d\n", Add(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y":
				if decision == "y" {
					continue
				}
				continue
			case "N":
				if decision == "n" {
					continue
				}
				return
			}
		case "2":
			fmt.Printf("Result = %d\n", Sub(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y":
				if decision == "y" {
					continue
				}
				continue
			case "N":
				if decision == "n" {
					continue
				}
				return
			}
		case "3":
			fmt.Printf("Result = %d\n", Mul(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y":
				if decision == "y" {
					continue
				}
				continue
			case "N":
				if decision == "n" {
					continue
				}
				return
			}
		case "4":
			fmt.Printf("Result = %f\n", Div(float64(num1), float64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y":
				if decision == "y" {
					continue
				}
				continue
			case "N":
				if decision == "n" {
					continue
				}
				return
			}
		case "5":
			fmt.Println("Commands......")
			fmt.Println("1  -->  Adds two numbers")
			fmt.Println("2  -->  Substract two numbers")
			fmt.Println("3  -->  Multiply two numbers")
			fmt.Println("4  -->  Divids two numbers")
			fmt.Println("5  -->  All the commands")
			fmt.Println("6  -->  Quits the program")
		case "6":
			return
		}
	}
}