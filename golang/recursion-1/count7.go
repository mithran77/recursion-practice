package main

import "fmt"

/*
https://codingbat.com/prob/p101409
Given a non-negative int n, return the count of the occurrences of 7 as a digit, so for example 717 yields 2. (no loops).
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6), while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).

count7(717) → 2
count7(7) → 1
count7(123) → 0
*/

func count7(num int) int {
	// Base case
	isSeven := 0
	if num < 1 {
		return 0
	}
	// Recursive case
	if num%10 == 7 {
		isSeven = 1
	}
	return isSeven + count7(num/10)
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("Number of sevens:", count7(num))
}
