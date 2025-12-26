# Sprout Language - Syntax

## Basic Syntax

### Variables
```python
sprout x = 10;              # Integer
sprout price = 19.99;       # Float  
sprout name = "Alice";      # String
sprout isValid = true;      # Boolean
```

### Operators
```python
# Arithmetic
+ - * / % **

# Comparison  
< > <= >= == !=

# Logical
&& || !  or  and or not
```

### Control Flow
```python
if (condition) {
    # code
} else {
    # code
}
```

### Print
```python
echo "Hello";
echo x + y;
```

### Comments
```python
# Single line comment
```

## REPL Commands

```
help       - Show help
env        - Show variables
clear      - Clear environment
trace on   - Enable trace mode
trace off  - Disable trace mode
exit/quit  - Exit REPL
```

## CLI Usage

```bash
# Interactive REPL
./sprout

# Run file
./sprun file.spr

# With trace
./sprun --trace file.spr

# With debug
./sprun --debug file.spr
```

## Error Messages

```
[Line X:Y] Parse error message
ERROR [Line X:Y]: Runtime error message
```

## Documentation

- `docs/USER_GUIDE.md` - Complete tutorial
- `docs/LEXER.md` - Lexer internals
- `docs/PARSER.md` - Parser details
- `docs/INTERPRETER.md` - Evaluator info
- `docs/API_REFERENCE.md` - Full API docs

## Examples

```python
# Calculate area
sprout pi = 3.14159;
sprout radius = 5.0;
sprout area = pi * radius ** 2.0;
echo area;

# Conditional
sprout age = 25;
if (age >= 18) {
    echo "Adult";
} else {
    echo "Minor";
}

# Complex expression
sprout result = (5 + 3) * 2 - 8 / 4;
echo result;  # 14
```

