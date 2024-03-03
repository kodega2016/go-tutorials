# Go Fundamentals

## Basic program structure

Like,other programming language each go program has a main function which is the entry point of the application.

```go
import "fmt"

package main

func main(){
  fmt.Println("Hello World")
}
```

## Variables

in go we can define the variables with the `var` keywords, we can define the variable with the multiple ways.

### Explicit Types

We can define the variable type explicitly with the var keywords.

```go
var name string="Software Development with Golang"
```

### Implicit Types

If we have not defined the variable type it will implicitly define its type on variable initialization.

```go
var age=12
```

here if we try to assign the string value to the age it will through the error because its type is `int` implicitly declared while the varible is assigned.

#### Short variable

we can declare variable with the short hand operators as well.so the basic sytax is as follows.

```go
username,role,is_active:="kodega2016","software developer",true
```

### Multi Types

another way to declare variables in Golang.

```go
var (
  company_name string="PortPro Nepal"
  company_size int=100
  company_address string="Sanepa,Nepal"
)
```

```

```
