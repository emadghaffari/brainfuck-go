package main

import (
	"fmt"
	"os"

	"github.com/emadghaffari/brainfuck-go/brain"
)

func main() {
	code, err := os.Open("file.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	brain := brain.New()
	brain.Compiler(code)
	brain.NewMachine(os.Stdin, os.Stdout)
	brain.Execute()
}
