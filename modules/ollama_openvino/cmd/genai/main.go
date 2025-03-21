package main

import (
	"fmt"
	"os"

	runner "github.com/ollama/ollama/genai/runner"
)

func main() {
	if err := runner.Execute(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}