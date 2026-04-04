# Go Varibles

## Objective
Today your going to be learning what a variable mean and the following below:
* Understand what a variable means
* Know the types of Go variable
* How to create a Go variable
* How to declare Go multiple variables
* How to professional name a variable

### Meaning of a Variable
Based on my understanding variable represents a value and it is also a container for storing values which can later be called.

### Types of Go variables
There are different types of variables
* `int` - stores positive and negative integers (numbers/digits) such as 123 and -123
* `float32`- stores fractional numbers and decimals e.g 21.7 or -21.7
* `string` - stores word e.g "Go Variables" and strings value are declared with double quotes.
* `bool` - this returns either true or false and stores it values. there are only two values with `bool` true of false.

Your going to learn more about Go variable types in `data_types.md`

### Creating (Declaring) Variables
In go there are two ways to declare a variable:

* using the `var` keyword
* using the `:=` operator

#### 1. Using the `var` Keyword:
You use the `var` keyword follwed by it's type and value:

##### Syntax
```go
var variablename type = value
```

This type of variable can be declared (created) outside a function main `func main() {}`:

##### Example:
```go
package main

import(
    "fmt"
)

var name string = "Chibueze"

func main() {
    fmt.Println(name)
}
```

##### Output:
```bash
✗ go run main.go

Chibueze
```
This tells use that the var keyword can be declare outside the function and inside.

##### Example:
```go
package main

import(
    "fmt"
)

var students int

func main() {
    student = 230

    var school string
    school = "Learn2earn"

    fmt.Println(student)
    fmt.Println(school)
}
```

##### Output:
```bash
✗ go run main.go

230
Learn2earn
```

#### 2. Using the `:=` Variable
This `:=` is called short variable declaration operator. It must be assigned a value for you to declare the variable. 

##### Syntax:
```
variablename := value
```
This short variable enables the computer decide the type based on the user value, the `:=` can only be declared inside a function

##### Example:
```go
package main

import(
    "fmt"
)

func main() {
    // declared in the function
    FullName := "Chibueze Charles Maxwell" //strings
    StudentID := 00019 // int
    account_Balance := 100000.07 // float32

    fmt.Println(FullName)
    fmt.Println(StudentID)
    fmt.Println(account_Balance)
}
```

##### Output:
```bash
✗ go run main.go

Chibueze Charles Maxwell
00019
100000.07
```
### Declaring Multiple Variables
In this section you're going to learn by practicing and following examples and personally modifying it while learning.

##### Example:
```go
package main

import(
    "fmt"
)

func main() {
    var a, b, c, d, e, int = 1, 4, 6, 9, 7

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
}
```

##### output:
```bash
✗ go run main.go

1
4
6
9
7
```
If the type of variable not stated you can declare aka create different types of variables on the same line.

##### Example:
```go
package main

import(
	"fmt"
)

func main() {
	var a, c = 7, "Billion"

	g, w := 21, "Trillion"

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(w)
}
```
##### output:
```bash
✗ go run main.go

7
Billion
21
Trillion
```

### Creating (Declaring) a Variable in a block
The declaring of variable enables readability in your code.

##### Example:
```go
package main

import(
	"fmt"
)

func main()  {
	var(
		a int
		b int = 7
		c string = "Seven"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```

##### output:
```bash
✗ go run main.go

0
7
Seven
```

### Go Variable Naming Rules
We have seen in our previous discussions that variable can have a short name (like a and b), so i'm going too give you the rules to name a Go variable professionally:

* A variable must start with a letter or an underscore character (_) e.g
```go
package main

import(
	"fmt"
)

func main()  {
	var _a string = "hello world!"
	fmt.Println(_a) // prints hello world!
}
```
* A variable cannot start with a digit e.g
```go
package main

import(
	"fmt"
)

func main()  {
	var 3_W string = "hello world!"
	fmt.Println(3_W)
}
```
##### output:
```bash
✗ go run main.go
# command-line-arguments
./main.go:8:6: syntax error: unexpected literal 3_, expected name
./main.go:8:7: '_' must separate successive digits
./main.go:9:15: '_' must separate successive digits
./main.go:9:16: syntax error: unexpected W in argument list; possibly missing comma or )
```
* A variable can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _)
```go
package main

import(
	"fmt"
)

func main()  {
	var M@well string = "hello world!"
	fmt.Println(M@well)
}
```

##### output:
```bash
✗ go run main.go
# command-line-arguments
./main.go:8:7: invalid character U+0040 '@'
./main.go:8:13: syntax error: unexpected string at end of statement
./main.go:9:15: invalid character U+0040 '@'
./main.go:9:16: syntax error: unexpected well in argument list; possibly missing comma or )
```
* Variables are case sensitive (name, Name and NAME are three different variables) e.g
```go
package main

import(
	"fmt"
)

func main()  {
	var fullName string = "hello world!"
	fmt.Println(FullName)
}
```
##### output:
```bash
✗ go run main.go
# command-line-arguments
./main.go:8:6: fullName declared and not used
./main.go:9:14: undefined: FullName
```

* There is no limit to the length of variable.
* A variables cannot contain space
* The variables cannot be any Go keywords e.g
```go
package main

import(
	"fmt"
)

func main()  {
	var var string = "hello world!"
	fmt.Println(var)
}
```
##### output:
```bash
✗ go run main.go
# command-line-arguments
./main.go:8:6: syntax error: unexpected var, expected name
./main.go:9:14: syntax error: unexpected var, expected expression
```

#### Multi-Word Variable Names
Variable names with more than one word can be difficult to read.

There are several techniques you can use to make them more readable:

##### Camel Case
The first letter starts with a lowercase after the first letter the remaining words starts with a capital letter:

>myVariableName = "Camel Case"

##### Pascal Case
Each word starts with a capital letter:

>MyVariableName = "Pascal Case"

##### Snake Case
Each word are separated with an underscore character:

>my_variable_name = "Snake Case"

#### That's it for variables, in our next topic open file `constants.md` were you will learn to use constants in your code.
