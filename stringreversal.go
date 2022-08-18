package main

func reversestring(s string) string {
	if s == "" {
		return ""
	}

	return reversestring(s[1:]) + string(s[0])
}
