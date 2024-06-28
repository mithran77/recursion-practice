package main

import "fmt"

/*
Given a string and a non-empty substring sub, compute recursively the number of times
that sub appears in the string, without the sub strings overlapping.

strCount("catcowcat", "cat") → 2
strCount("catcowcat", "cow") → 1
strCount("catcowcat", "dog") → 0
*/

func strCount(str string, sub string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if len(str)-i >= len(sub) && str[i:i+len(sub)] == sub {
			count++
			i += len(sub) - 1
		}
	}
	return count
}

func main() {
	fmt.Println(strCount("catcowcat", "cat"))
	fmt.Println(strCount("catcowcat", "cow"))
	fmt.Println(strCount("catcowcat", "dog"))
}
