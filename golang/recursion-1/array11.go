package main

import "fmt"

/*
https://codingbat.com/prob/p135988
Given an array of ints, compute recursively the number of times that the value 11 appears in the array.
We'll use the convention of considering only the part of the array that begins at the given index.
In this way, a recursive call can pass index+1 to move down the array. The initial call will pass in index as 0.

array11([1, 2, 11], 0) → 1
array11([11, 11], 0) → 2
array11([1, 2, 3, 4], 0) → 0
*/

func array11(nums []int, idx int) int {
	// Base case
	if idx >= len(nums) {
		return 0
	}
	// Recursive cases
	if nums[idx] == 11 {
		return 1 + array11(nums, idx+1)
	}
	return array11(nums, idx+1)
}

func main() {
	// var text []byte
	// fmt.Println("Enter a string:")
	// fmt.Scan(&text)
	// changePi(text, 0)
	fmt.Println(array11([]int{1, 2, 11}, 0))
	fmt.Println(array11([]int{11, 11}, 0))
	fmt.Println(array11([]int{1, 2, 3, 4}, 0))
	// fmt.Println("Pi changed to 3.14's:", string(text))
}

/*
array11(nums, 0)
array11(nums, 1)
array11(nums, 2)
array11(nums, 0)

array11()
array11()
array11()
array11()
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
