package main

import "fmt"

/*
https://codingbat.com/prob/p104029

Given a string, return recursively a "cleaned" string where adjacent chars that are the same have been reduced to a single char.
So "yyzzza" yields "yza".

stringClean("yyzzza") → "yza"
stringClean("abbbcdd") → "abcd"
stringClean("Hello") → "Helo"
*/

func stringClean(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if len(str)-i > 1 && str[i] == str[i+1] {
			continue
		}
		res += string(str[i])
	}
	return res
}

func main() {
	fmt.Println(stringClean("yyzzza"))
	fmt.Println(stringClean("abbbcdd"))
	fmt.Println(stringClean("Hello"))
	fmt.Println(stringClean("a"))
}
