package main

import "fmt"

func main() {
	var x, n int

	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%d\n", &n)

	fmt.Println("Output:", exponentiation(x, n))
}

func exponentiation(base, n int) int {
	result := 1

	for n > 0 {
		if n%2 == 0 {
			base *= base
			n /= 2
		} else {
			result *= base
			n -= 1
		}
	}

	return result
}
