package main

import "fmt"

/*
https://codingbat.com/prob/p158175

Given a string, compute recursively a new string where identical chars that are adjacent in the original string are separated from each other by a "*".

pairStar("hello") → "hel*lo"
pairStar("xxyy") → "x*xy*y"
pairStar("aaaa") → "a*a*a*a"
*/

func pairStar(str string) string {
	// Base case
	if len(str) <= 1 {
		return str
	}
	// Recursive cases
	if str[0] == str[1] {
		return string(str[0]) + "*" + pairStar(str[1:])
	}
	return string(str[0]) + pairStar(str[1:])
}

func main() {
	fmt.Println(pairStar("hello"))
	fmt.Println(pairStar("xxyy"))
	fmt.Println(pairStar("aaaa"))
	fmt.Println(pairStar("a"))
}
