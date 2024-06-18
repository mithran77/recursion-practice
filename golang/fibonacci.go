package main

import "fmt"

/*
https://codingbat.com/prob/p120015
The fibonacci sequence is a famous bit of mathematics, and it happens to have a recursive definition.
The first two values in the sequence are 0 and 1 (essentially 2 base cases).
Each subsequent value is the sum of the previous two values, so the whole sequence is: 0, 1, 1, 2, 3, 5, 8, 13, 21 and so on.
Define a recursive fibonacci(n) method that returns the nth fibonacci number, with n=0 representing the start of the sequence.


fibonacci(0) → 0
fibonacci(1) → 1
fibonacci(2) → 1

*/

func fibonacci(num int) int {
	// Base cases
	if num < 1 {
		return 0
	}
	if num == 1 {
		return 1
	}

	// Recursive case
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("nth fibonacci number:", fibonacci(num))
}

/*
fibonacci(213)
fibonacci(21)
fibonacci(2)
fibonacci(0)
*/
