# Contents

- [Contents](#contents)
  - [Variables](#variables)
  - [Print Formatted Data](#print-formatted-data)
  - [Data Types](#data-types)
  - [Taking Input](#taking-input)
    - [Pointer](#pointer)
  - [Code](#code)

## Variables

> Benefits

- Update the value once and can be seen through out the program
- Make our code flexible.

## Print Formatted Data

```go
fmt.printf("message %s",myvar)
```

General Formatting Styles Listed below ⬇️

```text
%v => the value in a default format when printing structs, the plus flag (%+v) adds field names
%#v => a Go-syntax representation of the value
%T => a Go-syntax representation of the type of the value
%% => a literal percent sign; consumes no value
```

- it takes a **template string** that contains the text that needs to be formatted
- **annotation verbs(or placeholder)** that tells the fmt functions how to format the variable passed in.

More about placeholders [placeholder](https://pkg.go.dev/fmt)

## Data Types

assigning variabale in go

```go
var name string
name="xyz"
const mod int = 80
```

Need to define data type after naming the variables.

- string
- int
- float
- boolean
- array
- map

Types of int GO

```text
uint8 -> Unsigned 8-bit integers (0 to 255)

uint16 -> Unsigned 16-bit integers (0 to 65535)

uint32 -> Unsigned 32-bit integers (0 to 4294967295)

uint64 -> Unsigned 64-bit integers (0 to 18446744073709551615)

int8 -> Signed 8-bit integers (-128 to 127)

int16 -> Signed 16-bit integers (-32768 to 32767)

int32 -> Signed 32-bit integers (-2147483648 to 2147483647)

int64 -> Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
```

uint defines Positive and whole numbers.

Types of Float in go

```text
float32 -> IEEE-754 32-bit floating-point numbers

float64 -> IEEE-754 64-bit floating-point numbers

complex64 -> Complex numbers with float32 real and imaginary parts

complex128 -> Complex numbers with float64 real and imaginary parts
```

For creating var , it can be done without using var and type

```go
name := "demo"
```

> Data type can be used if needed , but this will not work
> This method also doesn't work with const

## Taking Input

```go
fmt.scan()
```

This syntax is used for taking input in go

other than this ```go Scan()``` function there two more input function available for taking input

```go
fmt.Scanln()
fmt.Scanf()
```

1. Scan() : Scan scans text read from standard input, storing successive space-separated values into successive arguments.
2. Scanln() : Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
3. Scanf() : scanf scans text read from r, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.

### Pointer

-> Points to memory address of another variable. Which is reference of actual value .

-> It is also called **special variable** in GOlang.

```go
fmt.Println("Hey 👋", &name) // this will print memory address of variable
```

this will display the memory address of the variable

Output

```markdown
Hey 👋 0x14000010290
```

```go
var name string

1) fmt.Scan(name)
2) fmt.Scan(&name)

3) fmt.Println("Hey 👋", name)
```

Output

```markdown
ezio
Hey 👋 ezio
```

So in above given snippet of code
The syntax in **line 1** is taking input ```go fmt.Scan(name)``` but its just taking input and passing value of the variable to **line3**  which is null.

But if we use **line2** syntax for taking input it will pass the memory address of the variable where the data is stored  to **line3**.

## Code

Try given code for more understanding

[var.go](./var.go)
