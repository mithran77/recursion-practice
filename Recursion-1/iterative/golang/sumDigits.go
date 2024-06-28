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
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(sumDigits(126))
	fmt.Println(sumDigits(49))
	fmt.Println(sumDigits(12))
}

/**/
