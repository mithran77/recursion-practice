package main

import (
	"fmt"
	"strings"
)

/*
https://codingbat.com/prob/p105722

Given a string, compute recursively a new string where all the lowercase 'x' chars have been moved to the end of the string.

endX("xxre") → "rexx"
endX("xxhixx") → "hixxxx"
endX("xhixhix") → "hihixxx"
*/

func endX(str string) string {
	res := ""
	xs := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'x' {
			xs++
		} else {
			res += string(str[i])
		}
	}
	res += strings.Repeat("x", xs)
	return res
}

func main() {
	fmt.Println(endX("xxre"))
	fmt.Println(endX("xxhixx"))
	fmt.Println(endX("xhixhix"))
	fmt.Println(endX("a"))
}
