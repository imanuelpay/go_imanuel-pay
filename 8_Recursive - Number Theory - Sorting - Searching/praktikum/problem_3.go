package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}

func MaxSequence(arr []int) int {
	var count = []int{}
	result := arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if result+arr[i+1] < 0 {
			result = 0
			continue
		}

		result += arr[i+1]
		count = append(count, result)
	}

	sort.Ints(count)

	return count[len(count)-1]
}
