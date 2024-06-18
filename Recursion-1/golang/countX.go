package main

import "fmt"

/*
https://codingbat.com/prob/p170371
Given a string, compute recursively (no loops) the number of lowercase 'x' chars in the string.


countX("xxhixx") → 4
countX("xhixhix") → 3
countX("hi") → 0
*/

func countX(text string) int {
	xs := 0
	// Base case
	if text == "" {
		return 0
	}
	// Recursive case
	if text[0] == 'x' {
		xs = 1
	}
	return xs + countX(text[1:])
}

func main() {
	var text string
	fmt.Println("Enter a string:")
	fmt.Scan(&text)
	fmt.Println("Number of X's:", countX(text))
}

/*
countX("xxhixx")
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
