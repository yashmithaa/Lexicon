package main

import (
	"bufio"
	"fmt"
	"io"
	"lexicon/src/ast"
	"lexicon/src/evaluator"
	"lexicon/src/lexer"
	"lexicon/src/logger"
	"lexicon/src/parser"
	"os"
	"strings"
)

const PROMPT = "sprout> "

func main() {
	env := evaluator.NewEnvironment()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Sprout Programming Language REPL!")
	fmt.Println("Type 'help' for commands, 'exit' to quit")
	fmt.Println()

	for {
		fmt.Print(PROMPT)

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		line = strings.TrimSpace(line)

		// check for exit command
		if line == "exit" || line == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		// skip empty lines
		if line == "" {
			continue
		}

		// check for special commands
		if line == "help" {
			printHelp()
			continue
		}

		if line == "clear" {
			env = evaluator.NewEnvironment()
			fmt.Println("Environment cleared!")
			continue
		}

		if line == "env" {
			printEnvironment(env)
			continue
		}

		if line == "trace on" {
			logger.EnableTrace()
			fmt.Println("Trace mode enabled!")
			continue
		}

		if line == "trace off" {
			logger.DisableTrace()
			fmt.Println("Trace mode disabled!")
			continue
		}

		// evaluate the input
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		// Check for parser errors
		if len(p.Errors()) > 0 {
			fmt.Println("Parser errors:")
			for _, err := range p.Errors() {
				fmt.Printf("  %s\n", err)
			}
			continue
		}

		if len(program.Statements) == 0 {
			continue
		}

		result := evaluator.Eval(program, env)

		if result != nil {
			if errObj, ok := result.(*evaluator.Error); ok {
				fmt.Printf("Error: %s\n", errObj.Message)
			} else if result.Type() != evaluator.NULL_OBJ {
				// Only print results for non-PrintStatements
				// Check if the last statement was a print/echo
				if len(program.Statements) > 0 {
					if _, isPrint := program.Statements[len(program.Statements)-1].(*ast.PrintStatement); !isPrint {
						fmt.Println(result.Inspect())
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Sprout REPL Commands:")
	fmt.Println("  help       - Show this help message")
	fmt.Println("  clear      - Clear all variables from environment")
	fmt.Println("  env        - Show all variables in current environment")
	fmt.Println("  trace on   - Enable trace execution mode")
	fmt.Println("  trace off  - Disable trace execution mode")
	fmt.Println("  exit/quit  - Exit the REPL")
	fmt.Println()
	fmt.Println("Language Features:")
	fmt.Println("  sprout x = 10;           - Declare variable")
	fmt.Println("  x = 20;                  - Reassign variable")
	fmt.Println("  echo \"Hello\";            - Print output")
	fmt.Println("  if (x > 5) { echo x; }   - Conditionals")
	fmt.Println("  5 + 3 * 2;               - Expressions")
	fmt.Println("  true && false;           - Logical operations")
}

func printEnvironment(env *evaluator.Environment) {
	store := env.GetStore()
	if len(store) == 0 {
		fmt.Println("No variables defined yet.")
		return
	}

	fmt.Println("Current environment variables:")
	for name, value := range store {
		fmt.Printf("  %s = %s\n", name, value.Inspect())
	}
}
