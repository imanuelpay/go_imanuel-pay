package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	println("Nilai min:", min)
	println("Nilai max:", max)
}

func getMinMax(number ...*int) (min, max int) {
	min = *number[0]
	max = *number[0]

	for i := 0; i < len(number); i++ {
		if *number[i] < min {
			min = *number[i]
		}
		if *number[i] > max {
			max = *number[i]
		}
	}

	return
}
