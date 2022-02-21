package main

import "fmt"

func main() {
	var r, t, lp float64
	const pi = 3.14

	// Input
	fmt.Scanf("%g\n", &t)
	fmt.Scanf("%g\n", &r)

	lp = (2 * pi * r) * (r + t)
	fmt.Println("Output:", lp)
}
