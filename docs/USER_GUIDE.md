# Sprout Language - User Guide

## Introduction

Welcome to Sprout! This guide will help you learn the Sprout programming language from basics to advanced features.

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd Lexicon
   ```

2. Build the project:
   ```bash
   go build -o sprout cmd/repl/main.go
   go build -o sprun cmd/demo/main.go
   ```

### Running Programs

**Interactive REPL:**
```bash
./sprout
```

**Execute a file:**
```bash
./sprun examples/examples.spr
```

**With trace mode:**
```bash
./sprun --trace examples/examples.spr
```

**With debug logging:**
```bash
./sprun --debug examples/examples.spr
```

## Language Basics

### Comments

Use `#` for single-line comments:
```python
# This is a comment
sprout x = 10;  # This is also a comment
```

### Variables

**Declaration:**
```python
sprout x = 10;          # Integer
sprout pi = 3.14;       # Float
sprout name = "Alice";  # String
sprout flag = true;     # Boolean
```

**With type annotation (optional):**
```python
sprout age int = 25;
sprout price float = 19.99;
sprout message string = "Hello";
sprout isValid bool = false;
```

**Reassignment:**
```python
sprout x = 10;
x = 20;  # Reassign without 'sprout'
```

### Data Types

**Integers:**
```python
sprout a = 42;
sprout b = -10;
sprout c = 0;
```

**Floats:**
```python
sprout pi = 3.14159;
sprout e = 2.71828;
sprout negative = -5.5;
```

**Strings:**
```python
sprout greeting = "Hello, World!";
sprout quote = "She said \"Hello\"";  # Escape quotes
sprout multiline = "Line 1\nLine 2";   # Newlines
```

**Booleans:**
```python
sprout isTrue = true;
sprout isFalse = false;
```

### Printing Output

Use `echo` to print values:
```python
echo "Hello, Sprout!";
echo 42;
echo 3.14;
echo true;

sprout x = 10;
echo x;  # Prints: 10
```

## Operators

### Arithmetic Operators

```python
sprout a = 10;
sprout b = 3;

echo a + b;   # 13  - Addition
echo a - b;   # 7   - Subtraction
echo a * b;   # 30  - Multiplication
echo a / b;   # 3   - Division (integer division)
echo a % b;   # 1   - Modulo
echo 2 ** 8;  # 256 - Exponentiation
```

### Comparison Operators

```python
sprout x = 10;
sprout y = 20;

echo x < y;   # true  - Less than
echo x > y;   # false - Greater than
echo x <= y;  # true  - Less than or equal
echo x >= y;  # false - Greater than or equal
echo x == y;  # false - Equal to
echo x != y;  # true  - Not equal to
```

### Logical Operators

You can use symbols or words:
```python
# Using symbols
echo true && false;   # false - AND
echo true || false;   # true  - OR
echo !true;           # false - NOT

# Using words
echo true and false;  # false - AND
echo true or false;   # true  - OR
echo not true;        # false - NOT
```

### Operator Precedence

From highest to lowest:
1. `**` (Exponentiation)
2. `-x`, `!x`, `not x` (Unary minus, logical NOT)
3. `*`, `/`, `%` (Multiplication, division, modulo)
4. `+`, `-` (Addition, subtraction)
5. `<`, `>`, `<=`, `>=` (Comparisons)
6. `==`, `!=` (Equality)
7. `&&`, `and` (Logical AND)
8. `||`, `or` (Logical OR)

**Use parentheses to override precedence:**
```python
echo 2 + 3 * 4;       # 14
echo (2 + 3) * 4;     # 20

echo 2 ** 3 ** 2;     # 512 (right-associative: 2 ** 9)
echo (2 ** 3) ** 2;   # 64
```

## Expressions

### Simple Expressions

```python
sprout x = 5 + 3;          # 8
sprout y = x * 2;          # 16
sprout z = (y - 4) / 2;    # 6
```

### Complex Expressions

```python
sprout result = (10 + 5) * 2 - 8 / 4;  # 28

sprout a = 5;
sprout b = 10;
sprout c = a * b + (b - a) ** 2;  # 75
```

### String Concatenation

```python
sprout first = "Hello";
sprout last = "World";
sprout message = first + ", " + last + "!";
echo message;  # Hello, World!
```

### Mixed Type Arithmetic

```python
sprout int_val = 10;
sprout float_val = 3.14;
sprout result = int_val + float_val;  # 13.14 (automatic float conversion)
```

## Control Flow

### If-Else Statements

**Basic if:**
```python
sprout age = 20;

if (age >= 18) {
    echo "Adult";
}
```

**If-else:**
```python
sprout temperature = 25;

if (temperature > 30) {
    echo "Hot";
} else {
    echo "Cool";
}
```

**Nested conditions:**
```python
sprout score = 85;

if (score >= 90) {
    echo "Grade: A";
} else {
    if (score >= 80) {
        echo "Grade: B";
    } else {
        echo "Grade: C";
    }
}
```

### Complex Conditions

```python
sprout age = 25;
sprout hasLicense = true;

if (age >= 18 && hasLicense) {
    echo "Can drive";
} else {
    echo "Cannot drive";
}
```

```python
sprout x = 10;
sprout y = 20;
sprout z = 15;

if (x < y && y > z) {
    echo "y is the largest";
}
```

## Block Statements

Group multiple statements:
```python
sprout total = 0;

if (true) {
    sprout a = 10;
    sprout b = 20;
    total = a + b;
    echo total;
}
```

## REPL Commands

When using the interactive REPL:

- `help` - Show available commands
- `clear` - Clear all variables
- `env` - Show current variables
- `trace on` - Enable trace execution
- `trace off` - Disable trace execution
- `exit` or `quit` - Exit the REPL

## Common Patterns

### Counter Pattern

```python
sprout count = 0;
count = count + 1;
echo count;  # 1
```

### Swap Pattern (using temporary variable)

```python
sprout a = 10;
sprout b = 20;

sprout temp = a;
a = b;
b = temp;

echo a;  # 20
echo b;  # 10
```

### Conditional Assignment

```python
sprout x = 10;
sprout y = 20;
sprout max = 0;

if (x > y) {
    max = x;
} else {
    max = y;
}

echo max;  # 20
```

### Range Checking

```python
sprout value = 75;

if (value >= 0 && value <= 100) {
    echo "In range";
} else {
    echo "Out of range";
}
```

## Debugging

### Using Trace Mode

Enable trace to see execution flow:

**In REPL:**
```
sprout> trace on
Trace mode enabled!
sprout> sprout x = 10;
[TRACE] Eval: *ast.Program
[TRACE]   Eval: *ast.VariableDeclaration
[TRACE]     VariableDeclaration: x
[TRACE]       IntegerLiteral: 10
[TRACE]     Set variable x = 10
```

**In file execution:**
```bash
./sprun --trace myprogram.spr
```

### Using Debug Logging

```bash
./sprun --debug myprogram.spr
```

### Checking Variables

In REPL, use `env` to see all variables:
```
sprout> sprout x = 10;
sprout> sprout y = 20;
sprout> env
Current environment variables:
  x = 10
  y = 20
```

## Error Messages

Sprout provides detailed error messages with line numbers:

**Undefined variable:**
```
ERROR [Line 5:10]: identifier not found: x
```

**Type mismatch:**
```
ERROR: type mismatch: INTEGER + STRING
```

**Parse error:**
```
[Line 3:5] Expected next token to be =, got ; instead
```

**Division by zero:**
```
ERROR: division by zero
```

## Troubleshooting

**Problem:** Variable not found
**Solution:** Make sure you declared it with `sprout` first

**Problem:** Syntax error
**Solution:** Check for missing semicolons, parentheses, or braces

**Problem:** Type mismatch
**Solution:** Ensure you're using compatible types (can't add string to number)

**Problem:** Unexpected results
**Solution:** Use trace mode to see execution flow


