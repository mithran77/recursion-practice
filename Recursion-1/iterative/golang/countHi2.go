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
	count := 0
	for i := 0; i < len(str); i++ {
		if len(str)-i > 2 && str[i:i+3] == "xhi" {
			i += 2
		}
		if len(str)-i > 1 && str[i:i+2] == "hi" {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countHi2("ahixhi"))
	fmt.Println(countHi2("ahibhi"))
	fmt.Println(countHi2("xhixhi"))
	fmt.Println(countHi2("a"))
}
