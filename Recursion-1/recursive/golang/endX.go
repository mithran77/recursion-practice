package main

import "fmt"

/*
https://codingbat.com/prob/p105722

Given a string, compute recursively a new string where all the lowercase 'x' chars have been moved to the end of the string.

endX("xxre") → "rexx"
endX("xxhixx") → "hixxxx"
endX("xhixhix") → "hihixxx"
*/

func endX(str string) string {
	// Base case
	if str == "" {
		return str
	}
	// Recursive cases
	if str[0] == 'x' {
		return endX(str[1:]) + "x"
	} else {
		return string(str[0]) + endX(str[1:])
	}
}

func main() {
	fmt.Println(endX("xxre"))
	fmt.Println(endX("xxhixx"))
	fmt.Println(endX("xhixhix"))
	fmt.Println(endX("a"))
}
