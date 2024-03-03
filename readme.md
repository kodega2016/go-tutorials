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

### Short variable

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

## Conditionals

We can use if,else,if else for conditionals statement like other programming languages.

### If statement

```go
age:=10
if age<18{
  fmt.Println("you are not allowed to visit this site.")
}
```

### IF-Else statement

```go
age:=18
if age<10{
  fmt.Println("you are child")
}else if(age>10 && age<20){
  fmt.Println("you are a teenager")
}else{
  fmt.Println("you are a mature")
}
```

## Switch case

we can switch case for multiple conditionals checks.

```go
role:="developer"
switch(role){
case "admin":
  fmt.Println("you are an admin")
  break;
case "developer":
  fmt.Println("you are a developer")
  break;
default:
  fmt.Println("you are our guest")
}
```

In switch statement,there is a way t

```

```
