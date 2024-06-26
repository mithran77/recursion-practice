package main

import "fmt"

/*
https://codingbat.com/prob/p120015

The fibonacci sequence is a famous bit of mathematics, and it happens to have a recursive definition.
The first two values in the sequence are 0 and 1 (essentially 2 base cases).
Each subsequent value is the sum of the previous two values,
so the whole sequence is: 0, 1, 1, 2, 3, 5, 8, 13, 21 and so on.
Define a recursive fibonacci(n) method that returns the nth fibonacci number, with n=0 representing the start of the sequence.

fibonacci(0) → 0
fibonacci(1) → 1
fibonacci(2) → 1
*/

func fibonacci(n int) int {
	// Base case
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// Recursive cases
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(5))
}

/*
RT's
fibonacci(4)
fibonacci(3) + fibonacci(2)
fibonacci(2) + fibonacci(1) + fibonacci(1) + fibonacci(0)
fibonacci(1) + fibonacci(0) + fibonacci(1) + fibonacci(1) + fibonacci(0)

fibonacci(5)
fibonacci(4) + fibonacci(3)
fibonacci(3) + fibonacci(2) + fibonacci(2) + fibonacci(1)
fibonacci(2) + fibonacci(1) + fibonacci(1) + fibonacci(0) + fibonacci(1) + fibonacci(0) + fibonacci(1)
fibonacci(1) + fibonacci(0) fibonacci(1) + fibonacci(1) + fibonacci(0) + fibonacci(1) + fibonacci(0) + fibonacci(1)
*/
