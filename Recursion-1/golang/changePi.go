package main

import "fmt"

/*
Given a string, compute recursively (no loops) a new string where all appearances of "pi" have been replaced by "3.14".

changePi("xpix") → "x3.14x"
changePi("pipi") → "3.143.14"
changePi("pip") → "3.14p"
*/

// func changePi(text []byte, idx int) {
// 	// Base case
// 	if idx > len(text)-2 {
// 		fmt.Println("Pi changed to 3.14's:", string(text))
// 		return
// 	}
// 	// Recursive case
// 	if text[idx] == 'p' && text[idx+1] == 'i' {
// 		right := append([]byte{}, text[idx+2:]...)
// 		left := append(text[:idx], '3', '.', '1', '4')
// 		text = append(left, right...)
// 		changePi(text, idx+4)
// 	} else {
// 		changePi(text, idx+1)
// 	}
// }

func changePi(text []byte) string {
	// Base case
	if string(text) == "" {
		return string(text)
	}

	// Recursive case
	if text[0] == 'p' && text[1] == 'i' {
		return
	}
	return text[0]
}

func main() {
	var text []byte
	fmt.Println("Enter a string:")
	fmt.Scan(&text)
	// changePi(text, 0)
	changePi(text)
	// fmt.Println("Pi changed to 3.14's:", string(text))
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
