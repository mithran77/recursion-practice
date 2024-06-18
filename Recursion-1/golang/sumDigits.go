package main

import "fmt"

/*
https://codingbat.com/prob/p163932
Given a non-negative int n, return the sum of its digits recursively (no loops).
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6), while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).


sumDigits(126) → 9
sumDigits(49) → 13
sumDigits(12) → 3
*/

func sumDigits(num int) int {
	// Base case
	if num < 1 {
		return 0
	}
	// Recursive case
	return num%10 + sumDigits(num/10)
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("Sum of digits:", sumDigits(num))
}

/*
sumDigits(213)
sumDigits(21)
sumDigits(2)
sumDigits(0)
*/
