package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	fmt.Println(" ════════════════════════════════════════════════")
	fmt.Println("SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	fmt.Println(" ════════════════════════════════════════════════")

	fmt.Print("C&C> ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	cmd := strings.Fields(input)

	if cmd[0] != "calc" {
		fmt.Println("Invalid command")
		return
	}

	switch cmd[1] {
	case "add":
		add1, _ := strconv.ParseInt(cmd[2], 10, 64)
		add2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", add1+add2)

	case "sub":
		sub1, _ := strconv.ParseInt(cmd[2], 10, 64)
		sub2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", sub1-sub2)
	case "mul":
		mul1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mul2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", mul1*mul2)
	case "div":
		div1, _ := strconv.ParseInt(cmd[2], 10, 64)
		div2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", div1/div2)
	case "mod":
		mod1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mod2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", mod1%mod2)
	case "pow":
		pow1, _ := strconv.ParseInt(cmd[2], 10, 64)
		pow2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result:", math.Pow(float64(pow1), float64(pow2)))

	}
}
