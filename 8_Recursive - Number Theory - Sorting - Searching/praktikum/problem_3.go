package main

import (
	"fmt"
	"math"
)

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}

func primaSegiEmpat(high, wide, start int) {
	var mapPrime = map[int]int{}
	index := 0
	for i := start + 1; i < 10000; i++ {
		counter := 0
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				counter++
			}
		}

		if counter == 0 {
			mapPrime[index] = i
			index++
		}

		if index == wide*high {
			break
		}
	}

	var array = [][]int{}
	index = 0
	for i := 0; i < wide; i++ {
		var prime = []int{}
		for j := 0; j < high; j++ {
			prime = append(prime, mapPrime[index])
			index++
		}

		array = append(array, prime)
	}

	result := 0
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			fmt.Print(array[i][j], "\t")
			result += array[i][j]
		}
		fmt.Println()
	}

	fmt.Println(result)
}
