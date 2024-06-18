package main

import "fmt"

func printFromN(till int, n int) {
	if n < till {
		return
	}
	fmt.Println(n)
	printFromN(till, n-1)
}

func main() {
	var n int
	fmt.Print("Enter number n: ")
	fmt.Scan(&n)
	printFromN(1, n)
}

// Recursion Tree
// printToN(1, 3)
// printToN(1, 2)
// printToN(1, 1)
// printToN(1, 0)
