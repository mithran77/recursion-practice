package main

import "fmt"

/*
https://codingbat.com/prob/p158175

Given a string, compute recursively a new string where identical chars that are adjacent in the original string are separated from each other by a "*".

pairStar("hello") → "hel*lo"
pairStar("xxyy") → "x*xy*y"
pairStar("aaaa") → "a*a*a*a"
*/

func pairStar(text string) string {
	// Base case
	if len(text) <= 1 {
		return string(text[0])
	}
	// Recursive cases
	if text[0] == text[1] {
		return string(text[0]) + "*" + pairStar(text[1:])
	}
	return string(text[0]) + pairStar(text[1:])
}

func main() {
	fmt.Println(pairStar("hello"))
	fmt.Println(pairStar("xxyy"))
	fmt.Println(pairStar("aaaa"))
	fmt.Println(pairStar("a"))
}
