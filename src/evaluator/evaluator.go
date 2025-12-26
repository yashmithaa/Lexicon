package evaluator

import (
	"fmt"
	"lexicon/src/ast"
	"math"
)

// main entry point for evaluation
func Eval(node ast.Node, env *Environment) Object {
	switch node := node.(type) {

	// program node
	case *ast.Program:
		return evalProgram(node, env)

	// statements
	case *ast.VariableDeclaration:
		return evalVariableDeclaration(node, env)

	case *ast.PrintStatement:
		return evalPrintStatement(node, env)

	case *ast.ExpressionStatement:
		if node.Expression == nil {
			return NULL
		}
		return Eval(node.Expression, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	// expressions
	case *ast.IntegerLiteral:
		return &Integer{Value: node.Value}

	case *ast.FloatLiteral:
		return &Float{Value: node.Value}

	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.StringLiteral:
		return &String{Value: node.Value}

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	}

	return NULL
}

// evaluates a program node
func evalProgram(program *ast.Program, env *Environment) Object {
	var result Object = NULL

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		if returnValue, ok := result.(*Error); ok {
			return returnValue
		}
	}

	return result
}

// evaluates a block statement
func evalBlockStatement(block *ast.BlockStatement, env *Environment) Object {
	var result Object = NULL

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil && result.Type() == ERROR_OBJ {
			return result
		}
	}

	return result
}

// evaluates variable declaration
func evalVariableDeclaration(node *ast.VariableDeclaration, env *Environment) Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	env.Set(node.Name.Value, val)
	return val
}

// evaluates print statement
func evalPrintStatement(node *ast.PrintStatement, env *Environment) Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	fmt.Println(val.Inspect())
	return val
}

// evaluates if-else expression
func evalIfExpression(ie *ast.IfExpression, env *Environment) Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

// evaluates identifier (variable lookup)
func evalIdentifier(node *ast.Identifier, env *Environment) Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: %s", node.Value)
	}
	return val
}

// evaluates prefix expressions (-, !, not)
func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusOperator(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

// evaluates infix expressions (+, -, *, /, %, **, ==, !=, <, >, <=, >=, &&, ||)
func evalInfixExpression(operator string, left, right Object) Object {
	switch {
	// integer arithmetic
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	// float arithmetic
	case left.Type() == FLOAT_OBJ || right.Type() == FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)

	// boolean operations
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)

	// logical operations
	case operator == "&&":
		return nativeBoolToBooleanObject(isTruthy(left) && isTruthy(right))
	case operator == "||":
		return nativeBoolToBooleanObject(isTruthy(left) || isTruthy(right))

	// string concatenation
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ && operator == "+":
		leftVal := left.(*String).Value
		rightVal := right.(*String).Value
		return &String{Value: leftVal + rightVal}

	// type mismatch
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evaluates integer infix expressions
func evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}
		return &Integer{Value: leftVal / rightVal}
	case "%":
		if rightVal == 0 {
			return newError("modulo by zero")
		}
		return &Integer{Value: leftVal % rightVal}
	case "**":
		return &Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evaluates float infix expressions
func evalFloatInfixExpression(operator string, left, right Object) Object {
	var leftVal, rightVal float64

	// convert to float if needed
	switch left.Type() {
	case FLOAT_OBJ:
		leftVal = left.(*Float).Value
	case INTEGER_OBJ:
		leftVal = float64(left.(*Integer).Value)
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	switch right.Type() {
	case FLOAT_OBJ:
		rightVal = right.(*Float).Value
	case INTEGER_OBJ:
		rightVal = float64(right.(*Integer).Value)
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}
		return &Float{Value: leftVal / rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evaluates bang operator (!)
func evalBangOperator(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// evaluates minus prefix operator (-)
func evalMinusOperator(right Object) Object {
	switch right.Type() {
	case INTEGER_OBJ:
		value := right.(*Integer).Value
		return &Integer{Value: -value}
	case FLOAT_OBJ:
		value := right.(*Float).Value
		return &Float{Value: -value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

// helper: determines if value is truthy
func isTruthy(obj Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// helper: converts native bool to Boolean object
func nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// helper: creates a new error object
func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

// helper: checks if object is an error
func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}
