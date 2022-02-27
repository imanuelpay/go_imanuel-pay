package main

import "fmt"

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}

var temp = map[int]int{}

func fibo(n int) int {
	if value, cek := temp[n]; cek {
		return value
	}

	if n <= 1 {
		temp[n] = n
	} else {
		temp[n] = fibo(n-1) + fibo(n-2)
	}

	return temp[n]
}
