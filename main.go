package main

import (
	"fmt"
	"regexp_io/pkg/engine"
)

const (
	input  = "input2.txt"
	output = "output.txt"
)

func main() {
	err := engine.Engine(input, output)
	if err != nil {
		fmt.Println(err)
	}
}
