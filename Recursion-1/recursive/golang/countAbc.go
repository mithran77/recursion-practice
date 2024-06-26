package main

import "fmt"

/*
https://codingbat.com/prob/p161124

Count recursively the total number of "abc" and "aba" substrings that appear in the given string.

countAbc("abc") → 1
countAbc("abcxxabc") → 2
countAbc("abaxxaba") → 2
*/

func countAbc(str string) int {
	// Base case
	if len(str) < 3 {
		return 0
	}
	// Recursive cases
	if str[0:3] == "aba" || str[0:3] == "abc" {
		return 1 + countAbc(str[1:])
	} else {
		return countAbc(str[1:])
	}
}

func main() {
	fmt.Println(countAbc("abc"))
	fmt.Println(countAbc("abcxxabc"))
	fmt.Println(countAbc("abaxxaba"))
	fmt.Println(countAbc("a"))
}
