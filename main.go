package main

import "fmt"

func main() {
	// sum of n numbers
	fmt.Println(count(5))
	// reverse string
	fmt.Println(reversestring("hello"))
	// is palindrome?
	fmt.Println(palindrome("racecar"))
	// decimal to binary
	fmt.Println(decimaltobinary(233, ""))
	// binary search an array
	nums := []int{2, 5, 7, 11}
	fmt.Println(binarysearch(nums, 12, 0, len(nums)))
	//fibonacci
	fmt.Println(fibo(10))
	// merge sort
	sortthese := []int{2, 3, 1, 0, -1}
	mergesort(sortthese, 0, len(sortthese)-1)
	fmt.Println(sortthese)
}
