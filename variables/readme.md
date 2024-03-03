# Variables in go

We can define variable in go with the `var` keyword. mainly we can define variable with the following sytanx.

```go
var name string="Software Development with Golang"
```

This is the explicit declaration of the variable with explicit type `string`,
Also, we can define the variable without the type and it will assigned the type implictly.

```
go
var age=10
```

it will assign the type implictly, so its type is `int`
also we can declare variable with short hands

```go
username,role,age:="kodega2016","software developer",27
```

Another way to declare variable with multiple type with the following sytanx.

```go
var (
  company_name string="PortPro Nepal"
  company_size int=100
  is_active bool=true
)
```
