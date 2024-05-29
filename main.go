package main

import "fmt"

func main() {
	fmt.Println(invertText("hello world"))
}

func invertText(text string) string {
	var invertedText string
	for i := len(text) - 1; i >= 0; i-- {
		invertedText += string(text[i])
	}
	return invertedText
}
