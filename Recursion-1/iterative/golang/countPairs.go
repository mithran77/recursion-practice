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
	pairs := 0
	for i := 0; i < len(str); i++ {
		if len(str)-i > 2 && str[i] != str[i+1] && str[i] == str[i+2] {
			pairs++
		}
	}
	return pairs
}

func main() {
	fmt.Println(countPairs("axa"))
	fmt.Println(countPairs("axax"))
	fmt.Println(countPairs("axbx"))
	fmt.Println(countPairs("a"))
}
