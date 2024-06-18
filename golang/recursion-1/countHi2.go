package main

import "fmt"

/*
https://codingbat.com/prob/p143900

Given a string, compute recursively the number of times lowercase "hi" appears in the string,
however do not count "hi" that have an 'x' immedately before them.

countHi2("ahixhi") → 1
countHi2("ahibhi") → 2
countHi2("xhixhi") → 0
*/

func countHi2(str string) int {
	// Base case
	if len(str) < 2 {
		return 0
	}
	// Recursive cases
	if len(str) > 2 && str[0:3] == "xhi" {
		return countHi2(str[3:])
	} else if str[0:2] == "hi" {
		return 1 + countHi2(str[2:])
	} else {
		return countHi2(str[1:])
	}
}

func main() {
	fmt.Println(countHi2("ahixhi"))
	fmt.Println(countHi2("ahibhi"))
	fmt.Println(countHi2("xhixhi"))
	fmt.Println(countHi2("a"))
}
