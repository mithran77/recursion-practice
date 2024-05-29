package main

import "fmt"

func printToN(counter int, n int) {
	if counter > n {
		return
	}
	fmt.Println(counter)
	printToN(counter+1, n)
}

func main() {
	var n int
	fmt.Print("Enter number n: ")
	fmt.Scan(&n)
	printToN(1, n)
}

// Recursion Tree
// printToN(1, 3)
// printToN(2, 3)
// printToN(3, 3)
// printToN(4, 3)
