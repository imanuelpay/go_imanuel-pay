package main

import "fmt"

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}

func primeNumber(number int) bool {
	var counter = 0

	for i := 1; i <= number; i++ {
		if (number % i) == 0 {
			counter += 1
		} else {
			counter += 0
		}
	}

	if counter > 2 || number < 2 {
		return false
	} else {
		return true
	}
}
