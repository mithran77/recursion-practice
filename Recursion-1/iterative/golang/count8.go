package main

import "fmt"

/*
https://codingbat.com/prob/p192383

Given a non-negative int n, compute recursively (no loops) the count of the occurrences of 8 as a digit,
except that an 8 with another 8 immediately to its left counts double, so 8818 yields 4.
Note that mod (%) by 10 yields the rightmost digit (126 % 10 is 6),
while divide (/) by 10 removes the rightmost digit (126 / 10 is 12).

count8(8) → 1
count8(818) → 2
count8(8818) → 4
*/

func count8(n int) int {
	count := 0
	for n > 0 {
		if n%10 == 8 {
			if n > 9 && ((n/10)%10 == 8) {
				count += 2
			} else {
				count += 1
			}
		}
		n = n / 10
	}
	return count
}

func main() {
	fmt.Println(count8(8))
	fmt.Println(count8(818))
	fmt.Println(count8(8818))
}

/**/
