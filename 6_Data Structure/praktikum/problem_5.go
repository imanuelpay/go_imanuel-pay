package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output:", PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println("Output:", PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println("Output:", PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println("Output:", PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println("Output:", PairSum([]int{1, 5, 6, 7}, 6))
}

func PairSum(number []int, target int) []int {
	var index = map[int]int{}
	var a, b int
	var x1, x2 bool

	for i := 0; i < len(number); i++ {
		index[number[i]] = i
		if number[i] == target-number[i] {
			continue
		}

		a, x1 = index[number[i]]
		b, x2 = index[target-number[i]]
		if x1 && x2 {
			return []int{b, a}
		}
	}

	return []int{}
}
