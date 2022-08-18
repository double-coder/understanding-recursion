package main

import "fmt"

func decimaltobinary(num int, res string) string {
	if num == 0 {
		return res
	}

	res = fmt.Sprintf("%v", num%2) + res
	return decimaltobinary(num/2, res)
}
