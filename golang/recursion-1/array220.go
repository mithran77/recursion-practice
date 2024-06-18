package main

import "fmt"

/*
https://codingbat.com/prob/p173469
Given an array of ints, compute recursively if the array contains somewhere a value followed in the array by that value times 10.
We'll use the convention of considering only the part of the array that begins at the given index.
In this way, a recursive call can pass index+1 to move down the array.
The initial call will pass in index as 0.

array220([1, 2, 20], 0) → true
array220([3, 30], 0) → true
array220([3], 0) → false
*/

//	func array220(nums []int, idx int) bool {
//		// Base case
//		if idx >= len(nums) {
//			return false
//		}
//		// Recursive cases
//		for j := idx - 1; j >= 0; j-- {
//			if (nums[idx] / 10) == nums[j] {
//				return true
//			}
//		}
//		return array220(nums, idx+1)
//	}

func array220(nums []int, idx int) bool {
	// Base case
	if idx >= len(nums)-1 {
		return false
	}
	// Recursive cases
	if nums[idx]*10 == nums[idx+1] {
		return true
	}
	return array220(nums, idx+1)
}

func main() {
	fmt.Println(array220([]int{11, 2, 20}, 0))
	fmt.Println(array220([]int{3, 30}, 0))
	fmt.Println(array220([]int{3}, 0))
}
