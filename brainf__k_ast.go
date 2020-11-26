package main

import "fmt"

type value interface{}

type cons struct {
	car value
	cdr value
}

func Parse(code string) (value, string) {
	if code == "" {
		return nil, ""
	}
	switch code[0] {
	case '+', '-', '>', '<', '.', ',':
		ast, code1 := Parse(code[1:])
		return cons{code[0], ast}, code1
	case '[':
		ast1, code1 := Parse(code[1:])
		ast2, code2 := Parse(code1[1:])
		return cons{ast1, ast2}, code2
	case ']':
		return nil, code
	default:
		return Parse(code[1:])
	}
}

func Execute(v cons, tape map[int]int, dp int) (map[int]int, int) {
	switch x := v.car.(type) {
	case int:
		switch x {
		case '+':
			tape[dp] += 1
		case '-':
			tape[dp] -= 1
		case '>':
			dp += 1
		case '<':
			dp -= 1
		case '.':
			fmt.Printf("%c", x)
		}
	case cons:
		for tape[dp] != 0 {
			tape, dp = Execute(x, tape, dp)
		}
	}
	return tape, dp
}

func main() {
	ast, code := Parse("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	if code != "" {
		panic(fmt.Sprintf("Failed to parse. code: %s", code))
	}

	Execute(ast, map[int]int{}, 0)
}

// Tested with go version go1.13.1 linux/amd64
