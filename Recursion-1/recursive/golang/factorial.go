package main

import "fmt"

/*
https://codingbat.com/prob/p154669

Given n of 1 or more, return the factorial of n,
which is n * (n-1) * (n-2) ... 1. Compute the result recursively (without loops).

factorial(1) → 1
factorial(2) → 2
factorial(3) → 6
*/

func factorial(n int) int {
	// Base case
	if n <= 1 {
		return 1
	}
	// Recursive cases
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(1))
	fmt.Println(factorial(2))
	fmt.Println(factorial(3))
}

/**/
