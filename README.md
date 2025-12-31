# extio

`extio` is a simple Go package that provides an `Input` function similar to Python's `input()`, making it easy to take user input from the terminal.

## Installation

```bash
go get github.com/hamada-coolcode/extio
```

## Usage

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
