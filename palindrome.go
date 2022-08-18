package main

func palindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return palindrome(s[1 : len(s)-1])
	}
	return false
}
