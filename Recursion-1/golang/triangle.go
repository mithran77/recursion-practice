package main

import "fmt"

/*
https://codingbat.com/prob/p194781
We have triangle made of blocks. The topmost row has 1 block, the next row down has 2 blocks, the next row has 3 blocks, and so on.
Compute recursively (no loops or multiplication) the total number of blocks in such a triangle with the given number of rows.
*/

func blocks(rows int) int {
	// Base case
	if rows < 1 {
		return 0
	}
	// Recursive case
	return rows + blocks(rows-1)
}

func main() {
	var rows int
	fmt.Println("Number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Number of blocks:", blocks(rows))
}
