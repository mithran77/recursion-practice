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
	count := 0
	for i := 0; i < len(str)-2; i++ {
		if str[i:i+3] == "aba" || str[i:i+3] == "abc" {
			count++
			i += 2
		}
	}
	return count
}

func main() {
	fmt.Println(countAbc("abc"))
	fmt.Println(countAbc("abcxxabc"))
	fmt.Println(countAbc("abaxxaba"))
	fmt.Println(countAbc("a"))
}
