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
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	fmt.Println(factorial(1))
	fmt.Println(factorial(2))
	fmt.Println(factorial(3))
}

/**/
