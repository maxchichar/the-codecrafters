# Go Syntax

## Objective
So today you're going to be learning about the go syntax but first we have to know what a **syntax** means below.

### What is a syntax ?
Based on my understanding a syntax are elements and pieces that makes up a code, It also deals with how a code is structured in a computer language.

>>> Now that we know what a syntax means, let move on to learn what a Go syntax means.

### What is a go syntax ?
Based of what i've learnt from w3school a go syntax consist of package declaration, import packages, functions, statement, and expressions that makes up a golang code.

#### Example:
```go
package main

import(
    "fmt"
)

func main() {
    fmt.Println("Go Syntax!")
}
```

### Go Compact Code (not recommended)
Below is a more compact go of the previous **example**:
```go
package main; import("fmt"); func main() { fmt.Println("Go compact code");}
```
### How to run a go code
* First you open you're IDE 'vscode' after golang has been installed locally in your laptop.
* Create a file by using this command below in your terminal:
```zsh
touch main.go
```
* To code in this file you type this command in your terminal:
```bash
code main.go
```
* When the file has been opened and you can code print "I love golang":
```go
package main

import(
    "fmt"
)

func main() {
    fmt.Println("I love golang")
}

```
* Save your file `"crtl + s"` on window and linux while on mac `"cmd + s"`, then run the code using:
```zsh
go run main.go
```

#### **In our next topic open the file `comments.md` you will learn how to comment your codes in Go**

