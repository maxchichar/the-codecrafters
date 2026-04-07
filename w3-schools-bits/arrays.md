# Go Arrays

## Objectives
* Understand the meaning of Go arrays
* How to declare a Go array
* Implementation of Go arrays in your code

### Meaning of a Go array?
instead of declaring multiple variable under the same type `arrays` handled it by storing of multiple varibable of the same type in one variable.

### Creating an Array
Arrays can be declare in two ways.

#### 1. Using the `var` keyword:
##### Syntax:
```go
var array_name = [length]datatype{value} // Defined length

//Or

var array_name= [...]datatype{value}// inferred length
```

#### 2. Using the `:=` sign:
##### Syntax:
```go
array_name := [length]datatype{value} // Defined length

// Or

array_name := [...]datatype{value} // inferred length
```

### Implementation in Code
#### Example 1:
```go
package main

import(
	"fmt"
)

func main() {
	var a = [3]int{9,4,7}
	b := [...]int{4,1,9}

	fmt.Println(a)
	fmt.Println(b)
}
```

#### Example 2:
```go
package main

import(
	"fmt"
)

func main()  {
	var dogs = [6]string{"Havanese", "Shih tzu", "Pembroke welsh corgi", "King charles spaniel", "Chihuahua", "Bull dog"}

	fmt.Println(dogs)
}
```
#### Example 3:
Calculate the length of an array.
```go
package main

import(
	"fmt"
)

func main()  {
	NumStudents := [...]int{3,4,2,5,2,1,4,3,8,9}

	fmt.Println(len(NumStudents)) //len() is to calculate the length
}
```

#### Example 4:
Access element of an array:
```go
package main

import(
	"fmt"
)

func main()  {
	NumStudents := [10]int{3,4,2,5,2,1,4,3,8,9}

	fmt.Println(NumStudents[0])
	fmt.Println(NumStudents[7])
	fmt.Println(NumStudents[3])
}
```

#### Example 5:
Change element of an array:
```go
package main

import(
	"fmt"
)

func main()  {
	var SportCars = [7]string{"Porsche 718", "Porsche 911", "Lotus emira", "Alphine A110", "BMW M3/M4", "Aston martin vantage", "BMW M2"}

	SportCars[5] = "Ferrari 296 speciale A"
	fmt.Println(SportCars)
}
```