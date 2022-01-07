package main

import "fmt"

func main() {

	var word string
	fmt.Print("Insert word: ")
	fmt.Scanln(&word)
	fmt.Println(palindrome(word))
}

func palindrome(word string) string {
	var reverse string
	for i := len(word) - 1; i >= 0; i-- {
		reverse += string(word[i])
	}
	if word == reverse {
		return "Palindrome"
	} else {
		return "No Palindrome"
	}
}
