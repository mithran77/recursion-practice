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
	count := 0
	for i := 0; i < len(str)-1; i++ {
		if str[i:i+2] == "hi" {
			count++
			i++
		}
	}
	return count
}

func main() {
	fmt.Println(countHi("xxhixx"))
	fmt.Println(countHi("xhixhix"))
	fmt.Println(countHi("hi"))
}

/**/
