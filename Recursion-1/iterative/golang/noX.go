package main

import "fmt"

/*
https://codingbat.com/prob/p118230

Given a string, compute recursively a new string where all the 'x' chars have been removed.

noX("xaxb") → "ab"
noX("abc") → "abc"
noX("xx") → ""
*/
func noX(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] != 'x' {
			res += string(str[i])
		}
	}
	return res
}

func main() {
	fmt.Println(noX("xaxb"))
	fmt.Println(noX("abc"))
	fmt.Println(noX("xx"))
}

/**/
