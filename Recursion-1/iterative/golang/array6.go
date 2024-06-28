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

func array6(nums []int, index int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 6 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(array6([]int{1, 6, 4}, 0))
	fmt.Println(array6([]int{1, 4}, 0))
	fmt.Println(array6([]int{6}, 0))
}

/**/
