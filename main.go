package main

import "fmt"

func main() {
	fmt.Println(count(5))
	fmt.Println(reversestring("hello"))
	fmt.Println(palindrome("racecar"))
	fmt.Println(decimaltobinary(233, ""))

	nums := []int{2, 5, 7, 11}
	fmt.Println(binarysearch(nums, 12, 0, len(nums)))
}
