package main

func Palindrome(s string) string {
	runes := []rune(s)
	var palindrome []rune
	for i := len(s) - 1; i >= 0; i-- {
		palindrome = append(palindrome, runes[i])
	}

	return string(palindrome)
}
