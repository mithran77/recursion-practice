package main

/*
https://codingbat.com/prob/p107330
We have bunnies standing in a line, numbered 1, 2, ... The odd bunnies (1, 3, ..) have the normal 2 ears.
The even bunnies (2, 4, ..) we'll say have 3 ears, because they each have a raised foot.
Recursively return the number of "ears" in the bunny line 1, 2, ... n (without loops or multiplication).

bunnyEars2(0) → 0
bunnyEars2(1) → 2
bunnyEars2(2) → 5
*/

import "fmt"

func bunnyEars2(bunnies int) int {
	var ears int
	// Base case
	if bunnies < 1 {
		return 0
	}

	// Recursive case
	if bunnies%2 == 0 {
		ears = 3
	} else {
		ears = 2
	}
	return ears + bunnyEars2(bunnies-1)
}

func main() {
	var bunnies int
	fmt.Println("Number of bunnies:")
	fmt.Scan(&bunnies)
	fmt.Println("Number of ears:", bunnyEars2(bunnies))
}
