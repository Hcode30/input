package logger

import (
	"bufio"
	"fmt"
	"os"
)

var Version string = "1.0"

// input function takes in a variadic parameter of strings and returns a slice of strings.
// It prompts the user to input values for each string in the parameter and stores them in the slice.
// If no values are inputted, it returns nil.
func Input(str ...string) []string {

	var Inputs []string
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for _, s := range str {
		fmt.Print(s)
		scanner.Scan()
		Inputs = append(Inputs, scanner.Text())
	}
	if len(Inputs) == 0 {
		return nil
	}
	return Inputs
}
