package main

import "fmt"

/*
https://codingbat.com/prob/p101409

Given a non-negative int n, return the count of the occurrences of 7 as a digit,
so for example 717 yields 2. (no loops).
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6),
while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).

count7(717) → 2
count7(7) → 1
count7(123) → 0
*/

func count7(n int) int {
	// Base case
	if n < 1 {
		return 0
	}
	// Recursive case
	if n%10 == 7 {
		return 1 + count7(n/10)
	}
	return count7(n / 10)
}

func main() {
	fmt.Println(count7(717))
	fmt.Println(count7(7))
	fmt.Println(count7(123))
}

/*
RT's

count7(717)
1 + count7(71)
1 + count7(7)
1 + 1
*/
