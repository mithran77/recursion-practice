package main

import "fmt"

/*
https://codingbat.com/prob/p194781
We have triangle made of blocks. The topmost row has 1 block, the next row down has 2 blocks, the next row has 3 blocks, and so on.
Compute recursively (no loops or multiplication) the total number of blocks in such a triangle with the given number of rows.

triangle(0) → 0
triangle(1) → 1
triangle(2) → 3
*/

func triangle(rows int) int {
	// Base case
	if rows < 1 {
		return 0
	}
	// Recursive case
	return rows + triangle(rows-1)
}

func main() {
	fmt.Println(triangle(0))
	fmt.Println(triangle(1))
	fmt.Println(triangle(2))
}
