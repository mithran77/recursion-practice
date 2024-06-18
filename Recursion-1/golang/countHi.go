package main

import "fmt"

/*
https://codingbat.com/prob/p184029
Given a string, compute recursively (no loops) the number of times lowercase "hi" appears in the string.


countHi("xxhixx") → 1
countHi("xhixhix") → 2
countHi("hi") → 1

*/

func countHi(text string) int {
	// Base case
	if len(text) < 2 {
		return 0
	}
	// Recursive case
	if text[0] == 'h' && text[1] == 'i' {
		return 1 + countHi(text[2:])
	}
	return 0 + countHi(text[1:])
}

func main() {
	var text string
	fmt.Println("Enter a string:")
	fmt.Scan(&text)
	fmt.Println("Number of Hi's:", countHi(text))
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
