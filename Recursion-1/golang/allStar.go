package main

import "fmt"

/*
https://codingbat.com/prob/p183394

Given a string, compute recursively a new string where all the adjacent chars are now separated by a "*".

allStar("hello") → "h*e*l*l*o"
allStar("abc") → "a*b*c"
allStar("ab") → "a*b"
*/

func allStar(text string) string {
	// Base case
	if len(text) <= 1 {
		return string(text[0])
	}
	// Recursive cases
	return string(text[0]) + "*" + allStar(text[1:])
}

func main() {
	fmt.Println(allStar("hello"))
	fmt.Println(allStar("abc"))
	fmt.Println(allStar("ab"))
	fmt.Println(allStar("a"))
}
