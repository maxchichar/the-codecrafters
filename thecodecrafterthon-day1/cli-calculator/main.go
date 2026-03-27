package main

import (
	"fmt"
	"strconv"
)

var option int
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

	return float64(a) / float64(b)
}

func main()  {
	fmt.Println("Welcome To The Cli Calculator")
	fmt.Println()

	for {
		// Selection of operation
		fmt.Println()
		fmt.Println("Select an Operation by the number")

		op1 := 1
		fmt.Println(op1, ". Addition")

		op2 := 2
		fmt.Println(op2, ". Subtraction")

		op3 := 3
		fmt.Println(op3, ". Multiplication")

		op4 := 4
		fmt.Println(op4, ". Division")

		op5 := 5
		fmt.Println(op5, ". Help")

		op6 := 6
		fmt.Println(op6, ". Quit")
		fmt.Println()
		
		fmt.Print("Option: ")
		fmt.Scanln(&option)
		fmt.Println()

		// Options
		switch option {
		case 1:
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

			fmt.Printf("Result = %d\n", Add(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y", "y":
				continue
			case "N", "n":
				return
			}
		case 2:
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

			fmt.Printf("Result = %d\n", Sub(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y", "y":
				continue
			case "N", "n":
				return
			}
		case 3:
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

			fmt.Printf("Result = %d\n", Mul(int64(num1), int64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y", "y":
				continue
			case "N", "n":
				return
			}
		case 4:
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

			fmt.Printf("Result = %f\n", Div(float64(num1), float64(num2)))
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y", "y":
				continue
			case "N", "n":
				return
			}
		case 5:
			fmt.Println("Commands......")
			fmt.Println("1  -->  Adds two numbers")
			fmt.Println("2  -->  Substract two numbers")
			fmt.Println("3  -->  Multiply two numbers")
			fmt.Println("4  -->  Divids two numbers")
			fmt.Println("5  -->  All the commands")
			fmt.Println("6  -->  Quits the program")
			fmt.Println()

			fmt.Println("Do you want to quit or continue")
			fmt.Println()
			fmt.Println("Enter Y to continue. Enter N to quit")
			fmt.Scanln(&decision)

			switch decision {
			case "Y", "y":
				continue
			case "N", "n":
				return
			}
		case 6:
			fmt.Println("See You Again")
			return
		}
	}
}