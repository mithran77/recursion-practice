package main

import "fmt"

/*
https://codingbat.com/prob/p101372
Given a string, compute recursively (no loops) a new string where all the lowercase 'x' chars have been changed to 'y' chars.


changeXY("codex") → "codey"
changeXY("xxhixx") → "yyhiyy"
changeXY("xhixhix") → "yhiyhiy"
*/

func changeXY(str string) string {
	// Base case
	if str == "" {
		return str
	}
	// Recursive case
	if str[0] == 'x' {
		return "y" + changeXY(str[1:])
	} else {
		return string(str[0]) + changeXY(str[1:])
	}
}

func main() {
	fmt.Println(changeXY("codex"))
	fmt.Println(changeXY("xxhixx"))
	fmt.Println(changeXY("xhixhix"))
}

/*
RT's
countX("xxhixx")
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
