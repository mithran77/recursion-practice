package main

import "fmt"

/*
https://codingbat.com/prob/p183394

Given a string, compute recursively a new string where all the adjacent chars are now separated by a "*".

allStar("hello") → "h*e*l*l*o"
allStar("abc") → "a*b*c"
allStar("ab") → "a*b"
*/

func allStar(str string) string {
	res := ""
	for i := 0; i < len(str)-1; i++ {
		res += string(str[i]) + "*"
	}
	res += string(str[len(str)-1])
	return res
}

func main() {
	fmt.Println(allStar("hello"))
	fmt.Println(allStar("abc"))
	fmt.Println(allStar("ab"))
	fmt.Println(allStar("a"))
}
