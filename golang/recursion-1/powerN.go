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

func powerN(num, power int) int {
	// Base case
	if power < 1 {
		return 1
	}

	// Recursive case
	return num * powerN(num, power-1)
}

func main() {
	var num, power int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("Enter the power:")
	fmt.Scan(&power)
	fmt.Println("nth powerN number:", powerN(num, power))
}

/*
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
