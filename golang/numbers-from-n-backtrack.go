package main

import "fmt"

func printFromN(i int, n int) {
	if i > n {
		return
	}
	printFromN(i+1, n)
	fmt.Println(i)
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
