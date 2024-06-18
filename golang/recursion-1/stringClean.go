package main

import "fmt"

/*
https://codingbat.com/prob/p104029

Given a string, return recursively a "cleaned" string where adjacent chars that are the same have been reduced to a single char.
So "yyzzza" yields "yza".

stringClean("yyzzza") → "yza"
stringClean("abbbcdd") → "abcd"
stringClean("Hello") → "Helo"
*/

func stringClean(text string) string {
	// Base case
	if len(text) < 2 {
		return text
	}
	// Recursive cases
	if text[0] == text[1] {
		return string(stringClean(text[1:]))
	} else {
		return string(text[0]) + stringClean(text[1:])
	}
}

func main() {
	fmt.Println(stringClean("yyzzza"))
	fmt.Println(stringClean("abbbcdd"))
	fmt.Println(stringClean("Hello"))
	fmt.Println(stringClean("a"))
}
