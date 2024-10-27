package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func Calc(expression string) (float64, error) {
	// Парсим выражение
	node, err := parser.ParseExpr(expression)
	if err != nil {
		return 0, err
	}

	// Вычисляем значение выражения
	result, err := eval(node)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// eval - вспомогательная функция для вычисления значения выражения
func eval(node ast.Node) (float64, error) {
	switch n := node.(type) {
	case *ast.BasicLit:
		return strconv.ParseFloat(n.Value, 64)
	case *ast.BinaryExpr:
		x, err := eval(n.X)
		if err != nil {
			return 0, err
		}
		y, err := eval(n.Y)
		if err != nil {
			return 0, err
		}

		switch n.Op {
		case token.ADD:
			return x + y, nil
		case token.SUB:
			return x - y, nil
		case token.MUL:
			return x * y, nil
		case token.QUO:
			if y == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return x / y, nil
		default:
			return 0, fmt.Errorf("unsupported operation")
		}
	default:
		return 0, fmt.Errorf("unsupported expression type")
	}
}

func main() {
	expr := "3 + 5 * 2 - 8"
	result, err := Calc(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
