# extio

`extio` is a simple Go package that provides an `Input` function similar to Python's `input()`, making it easy to take user input from the terminal. It also supports typed input via generics.

## Installation

```bash
go get github.com/hamada-coolcode/extio
```

## Usage

### Basic String Input

```go
package main

import (
	"fmt"
	"github.com/hamada-coolcode/extio"
)

func main() {
	name := extio.Input("What is your name? ")
	fmt.Printf("Hello, %s!\n", name)
}
```

### Typed Input

You can use `InputAs[T]` to get input converted to a specific type (e.g., `int`, `float64`, `bool`).

```go
package main

import (
	"fmt"
	"github.com/hamada-coolcode/extio"
)

func main() {
	// Input as Integer
	age, err := extio.InputAs[int]("How old are you? ")
	if err != nil {
		fmt.Println("Invalid number")
		return
	}
	fmt.Printf("You are %d years old.\n", age)

	// Input as Boolean
	likesGo, err := extio.InputAs[bool]("Do you like Go? (true/false) ")
	if err == nil && likesGo {
		fmt.Println("Awesome!")
	}
}
```
