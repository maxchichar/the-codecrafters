# Go Output Functions

## Objective
* Understanding the output functions
* Learn how to implement it in your code
* Understand formatting Verb

## Types Of Output Funtions
Go has three output functions
* `Print()`
* `Println()`
* `Printf()`

### The Print() Function
This print it's arguments with their default format.
#### Example:
```go
package main

import(
    "fmt"
)

func main() {
    var a, b string = "hello", "world"
	// var c = "golang"
	fmt.Print(a)
	fmt.Print(b)
	// fmt.Print(c, "\n")
}
```
#### Output:
```bash
✗ go run main.go
helloworld                
```
we use the `\n` to print the argument in a new line
#### Example:
```go
package main

import(
	"fmt"
)

func main() {
	var a, b string = "hello", "world"
	var c = "golang"
	fmt.Print(a, "\n")
	fmt.Print(b, "\n")
	fmt.Print(c, "\n")
}
```

#### Output:
```bash
✗ go run main.go
hello
world
golang
```

You can also use `Print()` to print multiple variables

#### Example:

```go
package main

import(
	"fmt"
)

func main() {
	var a, b string = "hello", "world"
	fmt.Print(a, "\n", b)
}
```

#### Output:
```bash
✗ go run main.go
hello
world
```

### The Println() Function
This prints on a newline and adds whitespace between the argument unlike the `Print()`

#### Example:
```go
package main

import(
	"fmt"
)

func main() {
	var a, b string = "hello", "world"
	fmt.Println(a, b)
}
```
#### Output:
```bash
✗ go run main.go
hello world
```

### The Printf() Function
This formats it argument based on the formmatting verb and then prints it. It does not create a newline, so it will use `\n` to create newline.

#### Example:
For instance we are going to be using two types of formatting verbs below:
* `%v` this is used to print the `value` of an argument
* `%T` this is used to print the `type` of the argument

```go
package main

import(
	"fmt"
)

func main() {
	var a = "hello"
	var b = 7

	fmt.Printf("The value of a = %v and type: %T\n", a, a)
	fmt.Printf("The value of b = %v and type: %T\n", b, b)
}
```

#### Output:
```bash
✗ go run main.go
The value of a = hello and type: string
The value of b = 7 and type: int
```

## Go Formatting Verbs
Go provides several formatting verbs that can be used with the `Printf()` functions.

### General Formatting Verbs
>>>>Verb<---------------------->Description

* >%v  --->   This prints the value in defualt format
* >%#v --->   This prints the value in Go-syntax format
* >%T  --->   This prints the type of the value
* >%%  --->   This prints the `%` sign

#### Example:
```go
package main

import(
	"fmt"
)

func main() {
	var a = 21.7
	var txt = "Output Function"

	// For Integers
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v%%\n", a)

	// For Strings
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
```
#### Output:
```bash
✗ go run main.go
21.7
21.7
float64
21.7%
Output Function
"Output Function"
string
```
### Integer Formatting Verb

>Verb----------------------Description

>%b -----> Base 2

>%d -----> Base 10

>%+d -----> Base 10 and always shows sign

>%o -----> Base8

>%O -----> Base 8, leading with 0o

>%x -----> Base 16, lowercase

>%X -----> Base 16, uppercase

>%#x -----> Base 16, with leading 0x

>%4d ------> Pad with spaces (width 4, right justified)

>%-4d -----> Pad with spaces (width 4, left justified)

>%04d -----> Pad with zero (width 4)

#### Example:
```go
package main

import(
	"fmt"
)

func main()  {
	var a = 21

	fmt.Printf("%b\n", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%+d\n", a)
	fmt.Printf("%o\n", a)
	fmt.Printf("%O\n", a)
	fmt.Printf("%x\n", a)
	fmt.Printf("%X\n", a)
	fmt.Printf("%#x\n", a)
	fmt.Printf("%4d\n", a)
	fmt.Printf("%-4d\n", a)
	fmt.Printf("%04d\n", a)
}
```

#### Output:
```bash
✗ go run main.go
10101
21
+21
25
0o25
15
15
0x15
  21
21  
0021
```

### String Formatting Verbs
>Verb----------------Description

>`%s` Prints the value as plain string

>`%q`	Prints the value as a double-quoted string

>`%8s`	Prints the value as plain string (width 8, right justified)

>`%-8s`	Prints the value as plain string (width 8, left justified)

>`%x`	Prints the value as hex dump of byte values

>`%x`Prints the value as hex dump with spaces