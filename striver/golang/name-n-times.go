package main

// https://www.youtube.com/watch?v=un6PLygfXrA&list=PLgUwDviBIf0rGlzIn_7rsaR2FQ5e6ZOL9&index=2

import "fmt"

func printNameNTimes(name string, times *int) {
	if *times == 0 {
		return
	}
	fmt.Println(name)
	*times--
	printNameNTimes(name, times)
}

func main() {
	var cnt int
	fmt.Print("Type a number: ")
	fmt.Scan(&cnt)
	printNameNTimes("Mithran", &cnt)
}

// Recursion Tree
// printNameNTimes("Mithran", 3)
// printNameNTimes("Mithran", 2)
// printNameNTimes("Mithran", 1)
// printNameNTimes("Mithran", 0)
