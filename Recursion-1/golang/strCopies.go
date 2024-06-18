package main

import "fmt"

/*
Given a string and a non-empty substring sub, compute recursively if at least n
copies of sub appear in the string somewhere, possibly with overlapping. N will be non-negative.

strCopies("catcowcat", "cat", 2) → true
strCopies("catcowcat", "cow", 2) → false
strCopies("catcowcat", "cow", 1) → true
*/

func strCopies(str string, sub string, n int) bool {
	// Base cases
	if n < 1 {
		return true
	}
	if len(str) < len(sub) {
		return false
	}
	// Recursive cases
	if str[0:len(sub)] == sub {
		return strCopies(str[len(sub):], sub, n-1)
	} else {
		return strCopies(str[1:], sub, n)
	}
}

func main() {
	fmt.Println(strCopies("catcowcat", "cat", 2))
	fmt.Println(strCopies("catcowcat", "cow", 2))
	fmt.Println(strCopies("catcowcat", "cow", 1))
}
