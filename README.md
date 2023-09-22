# simple Input In GO

Input is a Go package that provides a simple and efficient way to prompt users for input in the command line.
With a simple Input Function, developers can save time and effort by eliminating the need to deal with pointers or addresses, and by allowing them to input multiple values in a single line of code.
The package includes a single function, input, which takes one or more string arguments as prompts and returns a slice of strings containing the user's input values.
The function uses the bufio package to read input values from the standard input, and then returns the input values as a slice.
This code is a great starting point for developers who want to quickly and easily prompt users for input in their Go projects, without having to write multiple lines of code for each message or input assignment.

## Installation

To use the Input function in your Go project, run the following command:

`go get github.com/Hcode30/input`

## Usage

To use the function, import it in your Go code:

```go
import i "github.com/Hcode00/simpleInput"
```

# Getting a single input value

To get a single input value, call the input function with a single prompt and retrieve the first element of the resulting slice:

```go
name := i.Input("Enter your name: ")[0] // John
fmt.Printf("Hello %s.\n", name) // Hello John
```

# Getting multiple input values

To get multiple input values, call the input function with multiple prompts and retrieve the element

```go
inputs := i.Input("Enter your name: ", "Enter your age: ", "Enter your gender: ") // John 21 Male
name, age, gender := inputs[0], inputs[1], inputs[2] // name = "John" , age = "21" , gender = "Male"
fmt.Printf("Hello %s, you are %s years old and you are a %s.\n", name, age, gender) //Hello John, you are 21 years old and you are a Male.\n
```

# Getting input values in a loop

To get input values in a loop, call the input function with a prompt inside the loop and append the resulting values to a slice:

```go
var inputs []string
for i := 0; i < 3; i++ {
    inputs = append(inputs, i.Input(fmt.Sprintf("Enter value %d: ", i+1))[0])
    // Enter value 1: 12
    // Enter value 2: 25
    // Enter value 3: "Hello World"
}
fmt.Printf("You entered: %v\n", inputs) // You entered: [12 25 Hello World]
```

# Contributing

If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.

# License

This package is licensed under the MIT License.
