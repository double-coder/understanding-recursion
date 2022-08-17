package main

func count(n int) int {
	// base case = stop recursion (we can multiple base cases)
	if n == 1 {
		return n
	}
	// recurse = minimum amount of work
	return 1 + count(n-1)
}
