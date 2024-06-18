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
	// Base cases
	if len(str) < len(sub) {
		return 0
	}
	// Recursive cases
	if str[:len(sub)] == sub && str[len(str)-len(sub):len(str)] == sub {
		return len(str)
	} else if str[0:len(sub)] == sub {
		return strDist(str[:len(str)-1], sub)
	} else if str[len(str)-len(sub):len(str)] == sub {
		return strDist(str[1:], sub)
	} else {
		return strDist(str[1:len(str)-1], sub)
	}
}

func main() {
	fmt.Println(strDist("catcowcat", "cat"))
	fmt.Println(strDist("catcowcat", "cow"))
	fmt.Println(strDist("cccatcowcatxx", "cat"))
}
