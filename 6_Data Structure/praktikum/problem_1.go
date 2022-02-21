package main

import "fmt"

func main() {
	fmt.Println("Output:", primeNumber(1000000007))
	fmt.Println("Output:", primeNumber(1500450271))
	fmt.Println("Output:", primeNumber(1000000000))
	fmt.Println("Output:", primeNumber(10000000019))
	fmt.Println("Output:", primeNumber(10000000033))
}

func primeNumber(number int) bool {
	var ket bool
	if number == 2 || number == 3 || number == 5 || number == 7 {
		ket = true
	} else if number <= 1 || number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		ket = false
	} else {
		ket = true
	}

	return ket
}
