# Sprout REPL Usage Guide

## Starting the REPL

### Option 1: Using the build script
```bash
./run.sh
```

### Option 2: Build and run manually
```bash
go build -o sprout cmd/repl/main.go
./sprout
```

### Option 3: Run directly with Go
```bash
go run cmd/repl/main.go
```

## Interactive Session Example

```
Welcome to the Sprout Programming Language REPL!
Type 'help' for commands, 'exit' to quit

sprout> help
Sprout REPL Commands:
  help      - Show this help message
  clear     - Clear all variables from environment
  env       - Show all variables in current environment
  exit/quit - Exit the REPL

Language Features:
  sprout x = 10;           - Declare variable
  x = 20;                  - Reassign variable
  echo "Hello";            - Print output
  if (x > 5) { echo x; }   - Conditionals
  5 + 3 * 2;               - Expressions
  true && false;           - Logical operations

sprout> sprout x = 10;
sprout> sprout y = 20;
sprout> echo x + y;
30
sprout> env
Current environment variables:
  x = 10
  y = 20
sprout> if (x < y) { echo "x is smaller"; }
x is smaller
sprout> clear
Environment cleared!
sprout> env
No variables defined yet.
sprout> exit
Goodbye!
```

## Tips

1. **Semicolons are optional** at the end of statements in the REPL
2. **Multi-line input** is not yet supported - write statements on one line
3. **View results** - Expressions automatically print their results
4. **Use `env`** to see all your variables at any time
5. **Use `clear`** to start fresh without restarting

## Common Use Cases

### Calculator
```
sprout> 2 ** 10;
1024
sprout> 100 / 3;
33
sprout> 3.14159 * 5.0 ** 2.0;
78.539750
```

### Variable Storage
```
sprout> sprout tax_rate = 0.08;
sprout> sprout price = 100.0;
sprout> sprout total = price + (price * tax_rate);
sprout> echo total;
108.000000
```

### Testing Logic
```
sprout> (5 > 3) && (10 < 20);
true
sprout> !false || false;
true
```

### String Operations
```
sprout> sprout name = "Sprout";
sprout> sprout msg = "Hello, " + name + "!";
sprout> echo msg;
Hello, Sprout!
```

## Keyboard Shortcuts

- **Ctrl+C** - Exit the REPL (alternative to typing `exit`)
- **Ctrl+D** - Exit the REPL (EOF)
- **Arrow Up/Down** - Not yet implemented (command history coming soon!)

## Error Messages

The REPL provides helpful error messages:

```
sprout> undefined_var;
Error: identifier not found: undefined_var

sprout> 5 + "hello";
Error: type mismatch: INTEGER + STRING

sprout> 10 / 0;
Error: division by zero
```
