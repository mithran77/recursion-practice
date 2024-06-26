package main

import "fmt"

/*
https://codingbat.com/prob/p163932

Given a non-negative int n, return the sum of its digits recursively (no loops).
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6),
while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).

sumDigits(126) → 9
sumDigits(49) → 13
sumDigits(12) → 3
*/

func sumDigits(n int) int {
	// Base case
	if n < 1 {
		return 0
	}
	// Recursive case
	return n%10 + sumDigits(n/10)
}

func main() {
	fmt.Println(sumDigits(126))
	fmt.Println(sumDigits(49))
	fmt.Println(sumDigits(12))
}

/*
RT's
sumDigits(213)
sumDigits(21)
sumDigits(2)
sumDigits(0)
*/
