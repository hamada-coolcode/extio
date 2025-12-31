package main

import (
	"fmt"

	"github.com/hamada-coolcode/extio"
)

func main() {
	name := extio.Input("What is your name? ")
	fmt.Printf("Hello, %s!\n", name)

	age, err := extio.InputAs[int]("How old are you? ")
	if err != nil {
		fmt.Printf("Error reading age: %v\n", err)
	} else {
		fmt.Printf("Next year you will be %d.\n", age+1)
	}

	height, err := extio.InputAs[float64]("Height in meters? ")
	if err == nil {
		fmt.Printf("Height: %.2f\n", height)
	}
}
