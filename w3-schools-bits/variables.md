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
> go run main.go

> Chibueze
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
> go run main.go

> 230
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
> go run main.go

> Chibueze Charles Maxwell
  00019
  100000.07
```


