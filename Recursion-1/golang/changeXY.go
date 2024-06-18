package main

import "fmt"

/*
https://codingbat.com/prob/p101372
Given a string, compute recursively (no loops) a new string where all the lowercase 'x' chars have been changed to 'y' chars.


changeXY("codex") → "codey"
changeXY("xxhixx") → "yyhiyy"
changeXY("xhixhix") → "yhiyhiy"
*/

func changeXY(text []byte, idx int) {
	// Base case
	if idx >= len(text) {
		return
	}
	// Recursive case
	if text[idx] == 'x' {
		text[idx] = 'y'
	}
	changeXY(text, idx+1)
}

func main() {
	var text []byte
	fmt.Println("Enter a string:")
	fmt.Scan(&text)
	changeXY(text, 0)
	fmt.Println("X changed to Y's:", string(text))
}

/*
countX("xxhixx")
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
