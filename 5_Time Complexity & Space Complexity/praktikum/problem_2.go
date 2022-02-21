package main

import "fmt"

func main() {
	var x, n int

	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%d\n", &n)

	fmt.Println("Output:", exponentiation(x, n))
}

func exponentiation(base, n int) int {
	var result int = 1

	for i := 1; i <= n/2; i++ {
		result *= base
	}

	if n%2 == 0 {
		result *= result
	} else {
		result = base * (result * result)
	}

	return result
}
