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
	// Base case
	if n < 1 {
		return 1
	}

	// Recursive case
	return base * powerN(base, n-1)
}

func main() {
	fmt.Println(powerN(3, 1))
	fmt.Println(powerN(3, 2))
	fmt.Println(powerN(3, 3))
}

/*
RT's

powerN(3,1)
powerN(3,0)

powerN(3,2)
powerN(3,1)
powerN(3,0)

powerN(3,3)
powerN(3,2)
powerN(3,1)
powerN(3,0)
*/
