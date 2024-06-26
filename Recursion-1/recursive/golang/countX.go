package main

import "fmt"

/*
https://codingbat.com/prob/p170371
Given a string, compute recursively (no loops) the number of lowercase 'x' chars in the string.

countX("xxhixx") → 4
countX("xhixhix") → 3
countX("hi") → 0
*/

func countX(str string) int {
	// Base case
	if str == "" {
		return 0
	}
	// Recursive cases
	if str[0] == 'x' {
		return 1 + countX(str[1:])
	}
	return countX(str[1:])
}

func main() {
	fmt.Println(countX("xxhixx"))
	fmt.Println(countX("xhixhix"))
	fmt.Println(countX("hi"))
}

/*
RT's

countX("xxhixx")
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
