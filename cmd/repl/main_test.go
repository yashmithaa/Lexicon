package main

import (
	"lexicon/src/evaluator"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printHelp() panicked: %v", r)
		}
	}()
	printHelp()
}

func TestPrintEnvironment(t *testing.T) {
	env := evaluator.NewEnvironment()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printEnvironment() panicked: %v", r)
		}
	}()
	printEnvironment(env)

	env.Set("x", &evaluator.Integer{Value: 10})
	printEnvironment(env)
}

func TestEnvironmentOperations(t *testing.T) {
	env := evaluator.NewEnvironment()
	env.Set("x", &evaluator.Integer{Value: 42})
	val, ok := env.Get("x")

	if !ok {
		t.Error("Failed to get variable 'x' from environment")
	}
	if intVal, ok := val.(*evaluator.Integer); ok {
		if intVal.Value != 42 {
			t.Errorf("Variable 'x' = %d, want 42", intVal.Value)
		}
	}
}

func TestREPLConstants(t *testing.T) {
	if PROMPT != "sprout> " {
		t.Errorf("PROMPT = %q, want 'sprout> '", PROMPT)
	}
}
