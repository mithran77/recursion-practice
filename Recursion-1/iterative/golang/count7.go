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
	count := 0
	for n > 0 {
		if n%10 == 7 {
			count++
		}
		n = n / 10
	}
	return count
}

func main() {
	fmt.Println(count7(717))
	fmt.Println(count7(7))
	fmt.Println(count7(123))
}

/**/
