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
	result := 1

	for i := 1; i <= n/2; i++ {
		result *= x
	}

	if n%2 == 0 {
		result *= result
	} else {
		result = x * (result * result)
	}

	return result
}
