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
	ears := 0
	for i := 1; i <= bunnies; i++ {
		if i%2 == 1 {
			ears += 2
		} else {
			ears += 3
		}
	}
	return ears
}

func main() {
	fmt.Println(bunnyEars2(0))
	fmt.Println(bunnyEars2(1))
	fmt.Println(bunnyEars2(2))
}
