package main

import "fmt"

/*
https://codingbat.com/prob/p161124

Count recursively the total number of "abc" and "aba" substrings that appear in the given string.

countAbc("abc") → 1
countAbc("abcxxabc") → 2
countAbc("abaxxaba") → 2
*/

func countAbc(text string) int {
	// Base case
	if len(text) < 3 {
		return 0
	}
	// Recursive cases
	if text[0:3] == "aba" || text[0:3] == "abc" {
		return 1 + countAbc(text[1:])
	} else {
		return countAbc(text[1:])
	}
}

func main() {
	fmt.Println(countAbc("abc"))
	fmt.Println(countAbc("abcxxabc"))
	fmt.Println(countAbc("abaxxaba"))
	fmt.Println(countAbc("a"))
}
