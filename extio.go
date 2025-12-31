package extio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input prints the prompt to standard output and reads a line from standard input.
// It mimics Python's input() function by stripping the trailing newline.
func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\r\n")
}

// InputAs reads a line using Input and attempts to convert it to type T.
// It uses fmt.Sscan for conversion.
func InputAs[T any](prompt string) (T, error) {
	var value T
	input := Input(prompt)
	_, err := fmt.Sscan(input, &value)
	return value, err
}
