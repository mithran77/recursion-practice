package main

import "fmt"

/*
https://codingbat.com/prob/p192383
Given a non-negative int n, compute recursively (no loops) the count of the occurrences of 8 as a digit,
except that an 8 with another 8 immediately to its left counts double, so 8818 yields 4.
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6), while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).

count8(8) → 1
count8(818) → 2
count8(8818) → 4
*/

func count8(num int) int {
	eights := 0
	// Base case
	if num < 1 {
		return 0
	}
	// Recursive case
	if num%10 == 8 {
		eights = 1
		if (num/10)%10 == 8 {
			eights = 2
		}
	}
	return eights + count8(num/10)
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("Number of 8's:", count8(num))
}

/*
sumDigits(213)
sumDigits(21)
sumDigits(2)
sumDigits(0)
*/
