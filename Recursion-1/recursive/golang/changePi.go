package main

import "fmt"

/*
Given a string, compute recursively (no loops) a new string where all appearances of "pi" have been replaced by "3.14".

changePi("xpix") → "x3.14x"
changePi("pipi") → "3.143.14"
changePi("pip") → "3.14p"
*/

func changePi(str string) string {
	// Base case
	if len(str) < 2 {
		return str
	}

	// Recursive cases
	if str[:2] == "pi" {
		return "3.14" + changePi(str[2:])
	}
	return string(str[0]) + changePi(str[1:])
}

func main() {
	fmt.Println(changePi("xpix"))
	fmt.Println(changePi("pipi"))
	fmt.Println(changePi("pip"))

}

/**/
