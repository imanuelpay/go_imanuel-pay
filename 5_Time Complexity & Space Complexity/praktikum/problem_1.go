package main

import "fmt"

func main() {
	fmt.Println("Output:", prime(1000000007))
	fmt.Println("Output:", prime(1500450271))
}

func prime(number int) string {
	var ket string
	if number == 2 || number == 3 || number == 5 || number == 7 {
		ket = "Bilangan Prima"
	} else if number <= 1 || number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		ket = "Bukan Bilangan Prima"
	} else {
		ket = "Bilangan Prima"
	}

	return ket
}
