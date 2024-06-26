package main

import "fmt"

/*
https://codingbat.com/prob/p118230

Given a string, compute recursively a new string where all the 'x' chars have been removed.

noX("xaxb") → "ab"
noX("abc") → "abc"
noX("xx") → ""
*/
func noX(str string) string {
	// Base case
	if str == "" {
		return str
	}

	// Recursive cases
	if str[0] == 'x' {
		return noX(str[1:])
	}
	return string(str[0]) + noX(str[1:])
}

func main() {
	fmt.Println(noX("xaxb"))
	fmt.Println(noX("abc"))
	fmt.Println(noX("xx"))
}

/**/
