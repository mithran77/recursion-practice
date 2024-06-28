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
	for i := 0; i < len(str); i++ {
		if len(str)-i >= len(sub) && str[i:i+len(sub)] == sub {
			n--
		}
	}
	return n == 0
}

func main() {
	fmt.Println(strCopies("catcowcat", "cat", 2))
	fmt.Println(strCopies("catcowcat", "cow", 2))
	fmt.Println(strCopies("catcowcat", "cow", 1))
}
