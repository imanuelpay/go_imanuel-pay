package main

import "math"

func main() {
	println(primeX(1))
	println(primeX(5))
	println(primeX(8))
	println(primeX(9))
	println(primeX(10))
}

func primeX(number int) int {
	var mapPrime = map[int]int{}
	index := 1
	for i := 2; i < 100000; i++ {
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

		if index > number {
			break
		}
	}

	value, _ := mapPrime[number]
	return value
}
