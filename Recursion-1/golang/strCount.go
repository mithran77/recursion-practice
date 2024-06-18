package main

import "fmt"

/*
Given a string and a non-empty substring sub, compute recursively the number of times
that sub appears in the string, without the sub strings overlapping.

strCount("catcowcat", "cat") → 2
strCount("catcowcat", "cow") → 1
strCount("catcowcat", "dog") → 0
*/

func strCount(str string, sub string) int {
	// Base case
	if len(str) < len(sub) {
		return 0
	}
	// Recursive cases
	if str[0:len(sub)] == sub {
		return 1 + strCount(str[len(sub):], sub)
	} else {
		return strCount(str[1:], sub)
	}
}

func main() {
	fmt.Println(strCount("catcowcat", "cat"))
	fmt.Println(strCount("catcowcat", "cow"))
	fmt.Println(strCount("catcowcat", "dog"))
}
