package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"lexicon/src/evaluator"
	"lexicon/src/lexer"
	"lexicon/src/logger"
	"lexicon/src/parser"
	"os"
	"strings"
)

func main() {
	// Command-line flags
	traceMode := flag.Bool("trace", false, "Enable trace execution mode")
	debugMode := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: sprun [--trace] [--debug] <filename.spr>")
		os.Exit(1)
	}

	filename := args[0]
	if !strings.HasSuffix(filename, ".spr") {
		fmt.Println("Error: File must have .spr extension")
		os.Exit(1)
	}

	// Set logging modes
	if *traceMode {
		logger.EnableTrace()
	}
	if *debugMode {
		logger.SetLevel(logger.DEBUG)
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parser errors
	if len(p.Errors()) > 0 {
		fmt.Println("Parser errors detected:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		os.Exit(1)
	}

	env := evaluator.NewEnvironment()

	fmt.Println("=== Sprout Interpreter ===")
	fmt.Println("Executing:", filename)
	if *traceMode {
		fmt.Println("Trace mode: ON")
	}
	fmt.Println()

	result := evaluator.Eval(program, env)

	// Check for runtime errors
	if result != nil && result.Type() == evaluator.ERROR_OBJ {
		fmt.Printf("\n%s\n", result.Inspect())
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("=== Done ===")
}
