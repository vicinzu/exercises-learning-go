package main

import (
	"errors"
	"fmt"
	"strconv"
)

var opmap = map[string]func(int, int) (int, error){
	"+": func(i, j int) (int, error) {
		return i + j, nil
	},
	"-": func(i, j int) (int, error) {
		return i - j, nil
	},
	"*": func(i, j int) (int, error) {
		return i * j, nil
	},
	"/": func(i, j int) (int, error) {
		if j == 0 {
			if i == 0 {
				return 0, nil
			}
			return 0, errors.New("division by zero")
		}
		return i / j, nil
	},
}

func errorf(format string, args ...interface{}) {
	fmt.Printf("Error: "+format+"\n", args...)
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"S"},
	}

	for _, expr := range expressions {
		if len(expr) != 3 {
			errorf("Invalid expression '%s'", expr)
			continue
		}

		op, ok := opmap[expr[1]]
		if !ok {
			errorf("unsupported operator '%s'", expr[1])
			continue
		}

		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			errorf("invalid 1st operand '%s'", err)
			continue
		}

		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			errorf("invalid 2nd operand '%s'", err)
			continue
		}

		result, err := op(p1, p2)
		if err != nil {
			errorf("%s", err)
			continue
		}
		fmt.Println("Result: ", result)
	}
}
