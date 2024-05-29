package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPalindrome Checker")
		fmt.Println("1. What is a palindrome?")
		fmt.Println("2. Check if a word is a palindrome")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward.")
			fmt.Println("For example: 'radar', 'level', '12321', and 'A man, a plan, a canal, Panama!' are palindromes.")
		case "2":
			fmt.Print("Enter a word: ")
			word, _ := reader.ReadString('\n')
			word = strings.TrimSpace(word)

			if isPalindrome2(word) {
				fmt.Printf("Yes, %s it's a palindrome!\n", word)
			} else {
				fmt.Printf("No, %s it's not a palindrome.\n", word)
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
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
	return true
}

// invertText function is not used because isPalindrome2 directly compares characters for efficiency.
// Keeping it here for reference or potential future use.
/*
func invertText(text string) string {
	var invertedText string
	for i := len(text) - 1; i >= 0; i-- {
		invertedText += string(text[i])
	}
	return invertedText
}
*/

// isPalindrome function is not used because isPalindrome2 provides better performance.
// Keeping it here for reference or potential future use.
/*
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	return invertText(s) == s
}
*/
