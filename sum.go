package main

func count(n int) int {
	// base case = stop recursion (we can multiple base cases)
	if n == 1 {
		return 1
	}
	// recurse = minimum amount of work
	return n + count(n-1)
}
