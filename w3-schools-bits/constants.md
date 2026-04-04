# Go Constants

## Objective
* Understand a constant
* Know how to create a constant
* Know the types of constants
* Rules of naming a Go constant
* How to use constants in your code
* Know the differences between constants and variables

### Meaning of Constants?
This is when a variable have a fixed value that cannot be changed once declared. It uses the `const` keyword for you to declare a constant. constants ar unchangeable and readonly.

### Declaring a Constant

##### Syntax:
```go
const CONSTNAME type = value
```
And note the value of a constant must be assigned when you declare it.

##### Example:
```go
package main

import(
	"fmt"
)

func main()  {
	const PI = 3.14

	fmt.Println(PI)
}
```

##### Output:
```bash
✗ go run main.go   
3.14
```

### Constant Rules
* Constant rules are same with that of variables, you can readmore on that in `variables.md`.
* They are usually written in uppercase for identification and differentiation from variables
* They can be declared both inside and outside the function e.g
```go
package main

import(
	"fmt"
)
const NAME string = "Chibueze"

func main()  {
	const PI = 3.14

	fmt.Println(PI)
	fmt.Println(NAME)
}
```

##### Output:
```bash
✗ go run main.go
3.14
```

### Constants Types

There are two types of constants:
* Typed constants
* Untyped constants

#### Typed constants
This are clearly defined with a defined type.

##### Example;
```go
package main

import(
	"fmt"
)

func main()  {
	const PI float64= 3.14
	// const NAME string = "Chibueze"

	fmt.Println(PI)
	// fmt.Println(NAME)
}
```
#### Untyped constants
This is the opposite of a Typed constant and they are declared without a type.

##### Example:
```go
package main

import(
	"fmt"
)

func main()  {
	const PI = 3.14
	// const NAME string = "Chibueze"

	fmt.Println(PI)
	// fmt.Println(NAME)
}
```
### Declaring Multiple Constants
Same with the variable in which they are grouped together for readablilty.

##### Example:
```go
package main

import(
	"fmt"
)

func main()  {
	const(
		PI = 3.14
		NAME  = "Chibueze"
		AGE = 21
	)

	fmt.Println(PI)
	fmt.Println(NAME)
	fmt.Println(AGE)
}
```

##### Output:
```bash
✗ go run main.go
3.14
Chibueze
21
```

#### In our next topic `output_func.md` you'll be learning different ways to produce an output.