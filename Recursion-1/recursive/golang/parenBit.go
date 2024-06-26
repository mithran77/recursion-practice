package main

import "fmt"

/*
https://codingbat.com/prob/p137918

Given a string that contains a single pair of parenthesis, compute recursively a new string made of only of the parenthesis and their contents,
so "xyz(abc)123" yields "(abc)".

parenBit("xyz(abc)123") → "(abc)"
parenBit("x(hello)") → "(hello)"
parenBit("(xy)1") → "(xy)"
*/

func parenBit(str string) string {
	// Base case
	if str[0] == '(' && str[len(str)-1] == ')' {
		return str
	}
	// Recursive cases
	if str[0] == '(' {
		return parenBit(str[:len(str)-1])
	}
	if str[len(str)-1] == ')' {
		return parenBit(str[1:])
	}
	return parenBit(str[1 : len(str)-1])
}

func main() {
	fmt.Println(parenBit("xyz(abc)123"))
	fmt.Println(parenBit("x(hello)"))
	fmt.Println(parenBit("(xy)1"))
}
