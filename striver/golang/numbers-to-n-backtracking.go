package main

import "fmt"

func printToN(i int, n int) {
	if i < 1 {
		return
	}
	printToN(i-1, n)
	fmt.Println(i)
}

func main() {
	var n int
	fmt.Print("Enter number n: ")
	fmt.Scan(&n)
	printToN(n, n)
}

// Recursion Tree
// printToN(3, 3)
// printToN(2, 3)
// printToN(1, 3)
// printToN(0, 3)
