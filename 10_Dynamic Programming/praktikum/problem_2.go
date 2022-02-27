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

func fibo(n int) int {
	var temp = map[int]int{}
	for i := 0; i <= n; i++ {
		if i <= 1 {
			temp[i] = i
		} else {
			temp[i] = temp[i-1] + temp[i-2]
		}
	}

	return temp[n]
}
