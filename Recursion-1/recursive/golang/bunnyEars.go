package main

/*
https://codingbat.com/prob/p183649
We have a number of bunnies and each bunny has two big floppy ears.
We want to compute the total number of ears across all the bunnies recursively (without loops or multiplication).

bunnyEars(0) → 0
bunnyEars(1) → 2
bunnyEars(2) → 4
*/

import "fmt"

func bunnyEars(bunnies int) int {
	// Base cases
	if bunnies < 1 {
		return 0
	}
	// Recursive case
	return 2 + bunnyEars(bunnies-1)
}

func main() {
	fmt.Println(bunnyEars(0))
	fmt.Println(bunnyEars(1))
	fmt.Println(bunnyEars(2))
}
