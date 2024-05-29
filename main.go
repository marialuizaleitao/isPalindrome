package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("hello world"))
}

func invertText(text string) string {
	var invertedText string
	for i := len(text) - 1; i >= 0; i-- {
		invertedText += string(text[i])
	}
	return invertedText
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	return invertText(s) == s
}

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return false
}
