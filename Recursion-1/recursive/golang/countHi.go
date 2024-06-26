package main

import "fmt"

/*
https://codingbat.com/prob/p184029
Given a string, compute recursively (no loops) the number of times lowercase "hi" appears in the string.


countHi("xxhixx") → 1
countHi("xhixhix") → 2
countHi("hi") → 1

*/

func countHi(str string) int {
	// Base case
	if len(str) < 2 {
		return 0
	}
	// Recursive cases
	if str[:2] == "hi" {
		return 1 + countHi(str[2:])
	}
	return countHi(str[1:])
}

func main() {
	fmt.Println(countHi("xxhixx"))
	fmt.Println(countHi("xhixhix"))
	fmt.Println(countHi("hi"))
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
