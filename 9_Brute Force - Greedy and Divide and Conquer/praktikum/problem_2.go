package main

import "fmt"

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}

var pecahan = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}

func moneyCoins(money int) []int {
	result := []int{}
	for i := len(pecahan) - 1; i >= 0; i-- {
		for money >= pecahan[i] {
			money -= pecahan[i]
			result = append(result, pecahan[i])
		}
	}

	return result
}
