package main

import "fmt"

/*
Given a string, compute recursively (no loops) a new string where all appearances of "pi" have been replaced by "3.14".

changePi("xpix") → "x3.14x"
changePi("pipi") → "3.143.14"
changePi("pip") → "3.14p"
*/

func changePi(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if len(str)-i > 1 && str[i:i+2] == "pi" {
			res += "3.14"
			i++
		} else {
			res += string(str[i])
		}
	}
	return res
}

func main() {
	fmt.Println(changePi("xpix"))
	fmt.Println(changePi("pipi"))
	fmt.Println(changePi("pip"))

}

/**/
