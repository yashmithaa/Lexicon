package main

import (
	"fmt"
	"lexicon/src/evaluator"
	"lexicon/src/lexer"
	"lexicon/src/parser"
)

func main() {
	input := `
	sprout x int = 10;
	sprout y = 20;
	sprout sum = x + y;
	sprout diff = y - x;
	sprout prod = x * y;
	sprout quot = y / x;
	sprout mod = y % x;
	sprout pow = 2 ** 3;
	echo "Arithmetic";
	echo sum;
	echo diff;
	echo prod;
	echo quot;
	echo mod;
	echo pow;
	
	sprout a float = 5.5;
	sprout b = 2.5;
	sprout fsum = a + b;
	sprout fprod = a * b;
	sprout fpow = 2.0 ** 3.0;
	echo "Floats";
	echo fsum;
	echo fprod;
	echo fpow;
	
	sprout gt = x > y;
	sprout lt = x < y;
	sprout eq = x == x;
	sprout neq = x != y;
	sprout lte = x <= y;
	sprout gte = y >= x;
	echo "Comparisons";
	echo gt;
	echo lt;
	echo eq;
	echo neq;
	echo lte;
	echo gte;
	
	x = 30;
	echo "Reassigned x";
	echo x;
	
	if (x > 20) {
		echo "x is greater than 20";
		if (x < 40) {
			sprout nested = x + 10;
			echo "Nested x + 10";
			echo nested;
		}
	} else {
		echo "x is not greater than 20";
	}
	
	sprout complex = (x + y) * 2 - (a / 2.0) ** 2;
	echo "Complex expression";
	echo complex;
	
	sprout neg = -x;
	sprout bang = !!true;
	echo "Prefix neg";
	echo neg;
	echo "bang";
	echo bang;
	
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
