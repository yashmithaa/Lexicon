package main

import (
	"fmt"
	"lexicon/src/evaluator"
	"lexicon/src/lexer"
	"lexicon/src/parser"
)

func main() {
	input := `
	sprout x = 10;
	sprout y = 20;
	sprout sum = x + y;
	echo "Sum:";
	echo sum;
	
	if (x < y) {
		echo "x is less than y";
	}
	
	sprout greeting = "Hello, " + "Sprout!";
	echo greeting;
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := evaluator.NewEnvironment()

	fmt.Println("=== Sprout Interpreter Demo ===")
	fmt.Println()
	evaluator.Eval(program, env)
	fmt.Println()
	fmt.Println("=== Done ===")
}
