package main

import "fmt"

/*
Given a string, return true if it is a nesting of zero or more pairs of parenthesis,
like "(())" or "((()))". Suggestion: check the first and last chars, and then recur on what's inside them.

nestParen("(())") → true
nestParen("((()))") → true
nestParen("(((x))") → false
*/

func nestParen(str string) bool {
	// Base cases
	if len(str) == 0 {
		return true
	}
	if len(str) == 1 {
		return false
	}
	// Recursive cases
	if str[0] == '(' && str[len(str)-1] == ')' {
		return nestParen(str[1 : len(str)-1])
	} else {
		return false
	}
}

func main() {
	fmt.Println(nestParen("(())"))
	fmt.Println(nestParen("((()))"))
	fmt.Println(nestParen("(((x))"))
	fmt.Println(nestParen("(((xyz))"))
}
