package main

import "fmt"

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}

func pangkat(base, pangkat int) int {
	var hasil = 1

	for i := 0; i < pangkat; i++ {
		hasil *= base
	}

	return hasil
}
