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

func array11(nums []int, index int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 11 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(array11([]int{1, 2, 11}, 0))
	fmt.Println(array11([]int{11, 11}, 0))
	fmt.Println(array11([]int{1, 2, 3, 4}, 0))
}

/**/
