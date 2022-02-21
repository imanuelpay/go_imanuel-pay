package main

import "fmt"

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}

func palindrome(input string) bool {
	var reverse string = ""
	for i := (len(input) - 1); i >= 0; i-- {
		reverse += string(input[i])
	}

	if input == reverse {
		return true
	} else {
		return false
	}
}
