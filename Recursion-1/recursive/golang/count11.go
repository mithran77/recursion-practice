package main

import "fmt"

/*
https://codingbat.com/prob/p167015
Given a string, compute recursively (no loops) the number of "11" substrings in the string.
The "11" substrings should not overlap.

count11("11abc11") → 2
count11("abc11x11x11") → 3
count11("111") → 1
*/

func count11(str string) int {
	// Base case
	if len(str) < 2 {
		return 0
	}
	// Recursive cases
	if str[0:2] == "11" {
		return 1 + count11(str[2:])
	} else {
		return count11(str[1:])
	}
}

func main() {
	fmt.Println(count11("11abc11"))
	fmt.Println(count11("abc11x11x11"))
	fmt.Println(count11("111"))
	fmt.Println(count11("a"))
}
