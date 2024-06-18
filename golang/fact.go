package main

import "fmt"

func fact(n int) int {
	// Base case
	if n < 1 {
		return 1
	}
	// Recursive case
	return n * fact(n-1)
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println("Factorial is :", fact(num))
}

/*
RT
fact(3)
fact(2)
fact(1)
fact(0)
*/
