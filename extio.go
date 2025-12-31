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
