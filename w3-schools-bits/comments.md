# Go Comments

## Objective
Welcome to go comments, today we are going to be learning about golang comment but first we need to know what **comments** mean in general.

### Meaning of comments ?
A comment is an unexecutable line in a code which can be used in explaining codes and making them readable, also when testing codes an alternative code it is used to prevent the main code from clashing with the alternative code.

This applies to all programming language including **GO**

### Types of comments in golang
There are two type of comments in golang:

* Go single-line comments
* Go multi-line comments

#### Go Single-line Comments
The go single-line comments starts with two forward slashes `//`. And any text or code behind `//` get's ignored by the compiler. It is mostly for shorter comments.

##### Example:
```go
// This is a single-line comment
package main

import(
    "fmt"
)

// This is a single-line comment
func main() {
    fmt.Println("Go single-line comments") // This is a single-line comment
}

```

#### Go Multi-line Comments
The go multi-line comments start with `/*` and ends with `*/`. And any text or code between `/*` and `*/` gets ignorned by the compiler. It is mostly used for longer comments.

##### Example:
```go
package main

import(
    "fmt"
)

/* 
This a multi-line comment.
It can span multiple lines
and will be ignored by the Go compiler
*/
func main() {
    fmt.Println("Go multi-line comments")
}

```
#### Comment to prevent code execution
You can use comments to prevent code from executing, that same code can be saved to use as reference.

##### Example:
```go
package main

import(
    "fmt"
)

func main() {
    fmt.Println("This is how to comment to prevent excution below")
    // fmt.Println("The line of code above is a go statement. Note this line does not execute")
}

```

#### **Open our next topic in `variables.md` were we learn how to use variables in our code**


