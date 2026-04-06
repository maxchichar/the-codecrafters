# Data Types

## Objectives
* Understand what data types mean
* Understand the types of data types

### What is the Meaning of Go Data Types?
A data type specify's the size and type of a variable value.

### The Types of Go Data Types?
There are three types of Go data types.
* Numeric
* Bool
* String

#### Numeric
This represents the integer `int32` `int64` types, floating point `float32` `float64` values, and complex types.
The integer has the signed and unsigned; the signed `int` can accept both negative and positive integers while the unsigned `uint` can only accept positive integers.

#### Bool
This represent either true or false value.

#### String
This returns a string value e.g `"Data types"`

#### Example:
```go
package main

import(
    "fmt"
)

func main() {
    var a string = "Data types" // A string value
    var b int32 = -21 // Integer under Numeric datatypes
    var c float64 = 21.7 // Floating point under Numeric datatypes 
    var d bool = true  // Boolean value
    var e uint = 21

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
}
```

#### Output
```bash
✗ go run main.go    
Data types
-21
21.7
true
21
```

