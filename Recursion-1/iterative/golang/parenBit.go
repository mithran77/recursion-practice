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
	for i, j := 0, len(str)-1; i < j; {
		if str[i] == '(' && str[j] == ')' {
			return str[i : j+1]
		} else if str[i] == '(' {
			j--
		} else if str[j] == ')' {
			i++
		} else {
			i++
			j--
		}
	}
	return str
}

func main() {
	fmt.Println(parenBit("xyz(abc)123"))
	fmt.Println(parenBit("x(hello)"))
	fmt.Println(parenBit("(xy)1"))
}
