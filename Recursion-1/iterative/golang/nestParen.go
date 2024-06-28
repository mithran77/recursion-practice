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
	for i, j := 0, len(str)-1; i <= j; {
		if str[i] == '(' && str[j] == ')' {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(nestParen("(())"))
	fmt.Println(nestParen("((()))"))
	fmt.Println(nestParen("(((x))"))
	fmt.Println(nestParen("(((xyz))"))
	fmt.Println(nestParen("(((yz))"))
	fmt.Println(nestParen(""))
}
