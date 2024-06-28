package main

import "fmt"

/*
Given a string and a non-empty substring sub, compute recursively the
largest substring which starts and ends with sub and return its length.

strDist("catcowcat", "cat") → 9
strDist("catcowcat", "cow") → 3
strDist("cccatcowcatxx", "cat") → 9
*/

func strDist(str string, sub string) int {
	for i, j := 0, len(str); i <= j; {
		if str[i:i+len(sub)] == sub && str[j-len(sub):j] == sub {
			return j - i
		} else if str[i:i+len(sub)] == sub {
			i++
		} else if str[j-len(sub):j] == sub {
			j--
		} else {
			i++
			j--
		}
	}
	return 0
}

func main() {
	fmt.Println(strDist("catcowcat", "cat"))
	fmt.Println(strDist("catcowcat", "cow"))
	fmt.Println(strDist("cccatcowcatxx", "cat"))
}
