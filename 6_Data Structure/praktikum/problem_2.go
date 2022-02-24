package main

import "fmt"

func main() {
	fmt.Println("Output:", pow(2, 3))
	fmt.Println("Output:", pow(7, 2))
	fmt.Println("Output:", pow(10, 5))
	fmt.Println("Output:", pow(17, 6))
	fmt.Println("Output:", pow(5, 3))
}

func pow(x, n int) int {
	var result int = 1

	for n > 0 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			result *= x
			n -= 1
		}
	}

	return result
}
