package main

import "fmt"

/*
https://codingbat.com/prob/p108997
Given an array of ints, compute recursively if the array contains a 6.
We'll use the convention of considering only the part of the array that begins at the given index.
In this way, a recursive call can pass index+1 to move down the array.
The initial call will pass in index as 0.


array6([1, 6, 4], 0) → true
array6([1, 4], 0) → false
array6([6], 0) → true
*/

func array6(nums []int, idx int) bool {
	// Base cases
	if idx >= len(nums) {
		return false
	}
	if nums[idx] == 6 {
		return true
	}
	// Recursive case
	return array6(nums, idx+1)
}

func main() {
	// var text []byte
	// fmt.Println("Enter a string:")
	// fmt.Scan(&text)
	// changePi(text, 0)
	fmt.Println(array6([]int{1, 6, 4}, 0))
	fmt.Println(array6([]int{1, 4}, 0))
	fmt.Println(array6([]int{6}, 0))
	// fmt.Println("Pi changed to 3.14's:", string(text))
}

/*
countX("xxhixx")
countX("xhixx")
countX("hixx")
countX("ixx")
countX("xx")
countX("x")
countX("")
*/
