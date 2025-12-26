# Lexicon - Sprout Programming Language Interpreter 🌱

An interpreter for **Sprout**, a simple programming language

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

### Run with Trace Mode
```bash
./sprun --trace filename.spr
```

### Run with Debug Logging
```bash
./sprun --debug filename.spr
```

### Interactive Session
```
sprout> sprout x = 10;
sprout> sprout y = 20;
sprout> echo x + y;
30
sprout> trace on
Trace mode enabled!
sprout> help
```

## Features

- **Lexer** - Tokenizes Sprout source code with line/column tracking
- **Parser** - Builds Abstract Syntax Tree (AST) with error collection
- **Interpreter** - Executes Sprout programs with full error handling
- **REPL** - Interactive command-line interface with environment inspection
- **Error Reporting** - Detailed errors with line and column numbers
- **Trace Execution** - Step-by-step debugging mode
- **Logging System** - Internal state logging for debugging
- **Documentation** - Comprehensive guides and API reference

## Language Overview

### Variables & Types
```sprout
sprout x = 10              # Integer
sprout pi = 3.14           # Float
sprout name = "Sprout"     # String
sprout flag = true         # Boolean

# With type annotations
sprout age int = 25
sprout price float = 19.99
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
sprout x = 10;  # Inline comment
```

## REPL Commands
- `help` - Show help and language features
- `env` - Show all variables
- `clear` - Clear environment
- `trace on` - Enable trace execution
- `trace off` - Disable trace execution
- `exit` - Quit REPL

## Error Reporting

Sprout provides detailed error messages with line and column information:

```
[Line 5:10] Expected next token to be =, got ; instead
ERROR [Line 3:5]: identifier not found: myVariable
ERROR: division by zero
ERROR: type mismatch: INTEGER + STRING
```

## Debugging Features

### Trace Execution
See step-by-step execution of your code:

```bash
./sprun --trace examples/examples.spr
```

Output:
```
[TRACE] Eval: *ast.Program
[TRACE]   Eval: *ast.VariableDeclaration
[TRACE]     VariableDeclaration: x
[TRACE]       IntegerLiteral: 10
[TRACE]     Set variable x = 10
```

### Environment Inspection
In REPL, use `env` to see all variables:
```
sprout> env
Current environment variables:
  x = 10
  y = 20
```

## Documentation

Comprehensive documentation is available in the `/docs` directory:

- **[User Guide](docs/USER_GUIDE.md)** - Complete language docs
- **[Syntax](docs/USER_GUIDE.md)** - Quick language syntax guide
- **[REPL](docs/USER_GUIDE.md)** - guide for REPL

## Project Structure

```
Lexicon/
├── cmd/
│   ├── repl/          # Interactive REPL
│   └── demo/          # File executor
├── src/
│   ├── token/         # Token definitions
│   ├── lexer/         # Lexical analyzer
│   ├── parser/        # Parser and AST
│   ├── ast/           # AST node definitions
│   ├── evaluator/     # Interpreter
│   └── logger/        # Logging system
├── docs/              # Documentation
├── examples/          # Example programs
```

## Testing
```bash
go test ./src/lexer      # Lexer tests
go test ./src/parser     # Parser tests
go test ./src/evaluator  # Evaluator tests
go test ./...            # All tests
```
