package main

import "fmt"

// Parameterized
// https://www.youtube.com/watch?v=69ZCDFy-OUo&list=RDCMUCJskGeByzRRSvmOyZOz61ig&index=1
// https://www.reddit.com/r/developersIndia/comments/15ds19h/difficulty_in_learning_recursion/

func sumN(i, sum int) {
	if i < 1 {
		fmt.Println("Sum: ", sum)
		return
	}
	sumN(i-1, sum+i)
}

func main() {
	var n int
	fmt.Print("Enter number n: ")
	fmt.Scan(&n)
	sumN(n, 0)
}

// RT
// sumToN(3, 0)
// sumToN(2, 3)
// sumToN(1, 5)
// sumToN(0, 6)
