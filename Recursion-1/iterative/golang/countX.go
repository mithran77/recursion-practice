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
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'x' {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(countX("xxhixx"))
	fmt.Println(countX("xhixhix"))
	fmt.Println(countX("hi"))
}

/**/
