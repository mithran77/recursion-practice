package main

import "fmt"

/*
https://codingbat.com/prob/p105722

Given a string, compute recursively a new string where all the lowercase 'x' chars have been moved to the end of the string.

endX("xxre") → "rexx"
endX("xxhixx") → "hixxxx"
endX("xhixhix") → "hihixxx"
*/

func endX(text string) string {
	// Base case
	if text == "" {
		return text
	}
	// Recursive cases
	if text[0] == 'x' {
		return endX(text[1:]) + "x"
	} else {
		return string(text[0]) + endX(text[1:])
	}
}

func main() {
	fmt.Println(endX("xxre"))
	fmt.Println(endX("xxhixx"))
	fmt.Println(endX("xhixhix"))
	fmt.Println(endX("a"))
}
