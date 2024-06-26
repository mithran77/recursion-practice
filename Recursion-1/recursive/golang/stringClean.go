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

func stringClean(str string) string {
	// Base case
	if len(str) < 2 {
		return str
	}
	// Recursive cases
	if str[0] == str[1] {
		// return string(str[0]) + string(stringClean(str[2:]))
		// Tackle adding the char at the end of repetition, as opposed to beginning
		return string(stringClean(str[1:]))
	} else {
		return string(str[0]) + stringClean(str[1:])
	}
}

func main() {
	fmt.Println(stringClean("yyzzza"))
	fmt.Println(stringClean("abbbcdd"))
	fmt.Println(stringClean("Hello"))
	fmt.Println(stringClean("a"))
}
