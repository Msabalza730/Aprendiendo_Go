package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter the Number to check Palindrome = ")
	fmt.Scanln(&num)
	fmt.Println(palindroNumber(num))
}

func palindroNumber(num int) string {
	var reverse int

	for temp := num; temp > 0; temp = temp / 10 {
		remainder := temp % 10
		reverse = reverse*10 + remainder
	}

	if num == reverse {
		return "Palindrome"
	} else {
		return "No Palindrome"
	}
}
