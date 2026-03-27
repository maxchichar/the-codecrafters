package main

import(
	"fmt"
)

var option string
var num1 int64
var num2 int64
var decision string


func Add(a, b int64) int64{
	return a + b
}

func Sub(a, b int64) int64 {
	return a - b
}

func Mul(a, b int64) int64 {
	return a * b
}

func Div(a, b int64) int64 {
	for {
		if b == 0 {
			fmt.Println("division by zero error:")
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

	fmt.Println("Select an Operation by the number")
	op1 := "1. Addition"
	fmt.Println(op1)

	op2 := "2. Subtraction"
	fmt.Println(op2)

	op3 := "3. Multiplication"
	fmt.Println(op3)

	op4 := "4. Division"
	fmt.Println(op4)

	op5 := "5. Quit"
	fmt.Println(op5)

	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("Input the numbers you want to add")
		fmt.Scan(&num1, &num2)
		fmt.Printf("Result = %d", Add(num1, num2))
		fmt.Println()

		fmt.Println("Do you want to quit or continue")
		fmt.Println("Enter Y to continue. Enter N to quit")
		fmt.Scanln(&decision)

		switch decision {
		case "Y":
			continue
		case "N":
			return
		}
		
	}
}