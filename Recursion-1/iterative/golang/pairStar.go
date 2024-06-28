package main

import "fmt"

/*
https://codingbat.com/prob/p158175

Given a string, compute recursively a new string where identical chars that are adjacent in
the original string are separated from each other by a "*".

pairStar("hello") → "hel*lo"
pairStar("xxyy") → "x*xy*y"
pairStar("aaaa") → "a*a*a*a"
*/

func pairStar(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] == str[i-1] {
			res += "*" + string(str[i])
		} else {
			res += string(str[i])
		}
	}
	return res
}

func main() {
	fmt.Println(pairStar("hello"))
	fmt.Println(pairStar("xxyy"))
	fmt.Println(pairStar("aaaa"))
	fmt.Println(pairStar("a"))
}
