package main

import "fmt"

/*
https://codingbat.com/prob/p158888
Given base and n that are both 1 or more, compute recursively (no loops) the value of base to the n power,
so powerN(3, 2) is 9 (3 squared).

powerN(3, 1) → 3
powerN(3, 2) → 9
powerN(3, 3) → 27

*/

func powerN(base, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= base
	}
	return res
}

func main() {
	fmt.Println(powerN(3, 1))
	fmt.Println(powerN(3, 2))
	fmt.Println(powerN(3, 3))
}

/**/
