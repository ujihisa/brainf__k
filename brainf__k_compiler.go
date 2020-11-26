package main

import "fmt"

func compile(code string) {
	fmt.Println("dp := 0")
	fmt.Println("tape := map[int]int{}")
	for _, c := range code {
		switch c {
		case '+':
			fmt.Println("tape[dp] += 1")
		case '-':
			fmt.Println("tape[dp] -= 1")
		case '>':
			fmt.Println("dp += 1")
		case '<':
			fmt.Println("dp -= 1")
		case '[':
			fmt.Println("for tape[dp] != 0 {")
		case ']':
			fmt.Println("}")
		case '.':
			fmt.Println("fmt.Printf(\"%c\", tape[dp])")
		}
	}
}

func main() {
	compile("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
}
