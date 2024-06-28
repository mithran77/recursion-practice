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
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] == 'x' {
			res += "y"
		} else {
			res += string(str[i])
		}
	}
	return res
}

func main() {
	fmt.Println(changeXY("codex"))
	fmt.Println(changeXY("xxhixx"))
	fmt.Println(changeXY("xhixhix"))
}

/**/
