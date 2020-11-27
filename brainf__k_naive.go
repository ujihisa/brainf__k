package main

import "fmt"

func Run(code string) {
	tape := make(map[int]int8)
	dp := 0
	ip := 0

	for ip < len(code) {
		// fmt.Printf("%05d dp = %v, tape[dp] = %v, code[ip] = %c\n", ip, dp, tape[dp], code[ip])
		switch code[ip] {
		case '+':
			tape[dp] += 1
		case '-':
			tape[dp] -= 1
		case '>':
			dp++
		case '<':
			dp--
		case '.':
			fmt.Printf("%c", tape[dp])
			// fmt.Println(tape[dp])
		case ',':
			panic(", is not implemented yet")
		case '[':
			if tape[dp] == 0 {
				// fmt.Printf("nextMatchingClose: %v -> ", ip)
				ip = nextMatchingClose(code, ip)
				// fmt.Printf("%v\n", ip)
			}
		case ']':
			if tape[dp] != 0 {
				// fmt.Printf("prevMatchingOpen: %v -> ", ip)
				ip = prevMatchingOpen(code, ip)
				// fmt.Printf("%v\n", ip)
			}
		default:
			// comment
		}
		ip++
	}
}

func nextMatchingClose(code string, ip int) int {
	ip++
	switch code[ip] {
	case '[':
		return nextMatchingClose(code, nextMatchingClose(code, ip))
	case ']':
		return ip
	default:
		return nextMatchingClose(code, ip)
	}
}

func prevMatchingOpen(code string, ip int) int {
	ip--
	switch code[ip] {
	case ']':
		return prevMatchingOpen(code, prevMatchingOpen(code, ip))
	case '[':
		return ip
	default:
		return prevMatchingOpen(code, ip)
	}
}

func main() {
	Run("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
}

// Tested with go version go1.13.1 linux/amd64
