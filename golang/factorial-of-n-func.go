package main

import "fmt"

// Parameterized
// https://www.youtube.com/watch?v=69ZCDFy-OUo&list=RDCMUCJskGeByzRRSvmOyZOz61ig&index=1
// https://www.reddit.com/r/developersIndia/comments/15ds19h/difficulty_in_learning_recursion/

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var n int
	fmt.Print("Enter number n: ")
	fmt.Scan(&n)
	fmt.Println(fact(n))
}

// RT
// sumToN(3)
// sumToN(2)
// sumToN(1)
// sumToN(0)
