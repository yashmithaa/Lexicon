package main

import (
	"fmt"
	"io/ioutil"
	"lexicon/src/evaluator"
	"lexicon/src/lexer"
	"lexicon/src/parser"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sprun <filename.spr>")
		os.Exit(1)
	}

	filename := os.Args[1]
	if !strings.HasSuffix(filename, ".spr") {
		fmt.Println("Error: File must have .spr extension")
		os.Exit(1)
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()
	env := evaluator.NewEnvironment()

	fmt.Println("=== Sprout Interpreter ===")
	fmt.Println("Executing:", filename)
	fmt.Println()
	evaluator.Eval(program, env)
	fmt.Println()
	fmt.Println("=== Done ===")
}
