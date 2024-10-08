package main

import "fmt"

/*
https://codingbat.com/prob/p154048

We'll say that a "pair" in a string is two instances of a char separated by a char.
So "AxA" the A's make a pair. Pair's can overlap, so "AxAxA" contains 3 pairs -- 2 for A and 1 for x.
Recursively compute the number of pairs in the given string.

countPairs("axa") → 1
countPairs("axax") → 2
countPairs("axbx") → 1
*/

func countPairs(str string) int {
	// Base case
	if len(str) < 3 {
		return 0
	}
	// Recursive cases
	if str[0] == str[2] && str[0] != str[1] {
		return 1 + countPairs(str[1:])
	} else {
		return countPairs(str[1:])
	}
}

func main() {
	fmt.Println(countPairs("axa"))
	fmt.Println(countPairs("axax"))
	fmt.Println(countPairs("axbx"))
	fmt.Println(countPairs("a"))
}
