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
	count := 0
	for i := 0; i < len(str); i++ {
		if len(str)-i > 1 && str[i:i+2] == "11" {
			count++
			i++
		}
	}
	return count
}

func main() {
	fmt.Println(count11("11abc11"))
	fmt.Println(count11("abc11x11x11"))
	fmt.Println(count11("111"))
	fmt.Println(count11("a"))
}
