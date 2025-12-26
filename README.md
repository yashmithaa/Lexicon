# Lexicon - Sprout Programming Language Interpreter

An interpreter for **Sprout**, a simple programming language with a playful design.

## Quick Start

### Build the Interpreter
```bash
./run.sh
```

### Run REPL
```bash
./sprout
```

### Run Sprout Files
```bash
./sprun filename.spr
```

### Interactive Session
```
sprout> sprout x = 10;
sprout> sprout y = 20;
sprout> echo x + y;
30
sprout> help
```

## Features

- **Lexer** - Tokenizes Sprout source code
- **Parser** - Builds Abstract Syntax Tree (AST)
- **Interpreter** - Executes Sprout programs with full error handling
- **REPL** - Interactive command-line interface

## Language Overview

### Variables & Types
```sprout
sprout x = 10              # Integer
sprout pi = 3.14           # Float
sprout name = "Sprout"     # String
sprout flag = true         # Boolean
```

### Operators
- **Arithmetic:** `+` `-` `*` `/` `%` `**`
- **Comparison:** `<` `>` `<=` `>=` `==` `!=`
- **Logical:** `&&` `||` `!` (or `and` `or` `not`)

### Control Flow
```sprout
if (x > 5) {
    echo "Greater";
} else {
    echo "Smaller";
}
```

### Comments
```sprout
# This is a comment
```

## REPL Commands
- `help` - Show help
- `env` - Show all variables
- `clear` - Clear environment
- `exit` - Quit REPL

## Testing
```bash
go test ./...      # Run all tests (20/20 passing)
```
