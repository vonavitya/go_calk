package calc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	if tokens == nil {
		return 0, fmt.Errorf("invalid expression")
	}
	result, err := parseExpression(tokens)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(expression string) []string {
	var tokens []string
	var number strings.Builder

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			number.WriteRune(char)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			if char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')' {
				tokens = append(tokens, string(char))
			} else if !unicode.IsSpace(char) {
				return nil
			}
		}
	}
	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}
	return tokens
}

func parseExpression(tokens []string) (float64, error) {
	var stack []float64
	var ops []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	applyOp := func() error {
		if len(stack) < 2 || len(ops) == 0 {
			return fmt.Errorf("invalid expression")
		}
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		op := ops[len(ops)-1]
		stack = stack[:len(stack)-2]
		ops = ops[:len(ops)-1]

		var result float64
		switch op {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				return fmt.Errorf("division by zero")
			}
			result = a / b
		default:
			return fmt.Errorf("invalid operator")
		}
		stack = append(stack, result)
		return nil
	}

	for _, token := range tokens {
		if token == "(" {
			ops = append(ops, token)
		} else if token == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("mismatched parentheses")
			}
			ops = ops[:len(ops)-1]
		} else if precedence[token] > 0 {
			for len(ops) > 0 && precedence[ops[len(ops)-1]] >= precedence[token] {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, token)
		} else {
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack = append(stack, value)
		}
	}

	for len(ops) > 0 {
		if err := applyOp(); err != nil {
			return 0, err
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}
