package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}

func FindMinAndMax(arr []int) string {
	min := arr[0]
	max := arr[0]

	minString := "min: " + strconv.Itoa(min) + " index: 0"
	maxString := "max: " + strconv.Itoa(max) + " index: 0"

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minString = "min: " + strconv.Itoa(min) + " index: " + strconv.Itoa(i)
		}
		if arr[i] > max {
			max = arr[i]
			maxString = ", max: " + strconv.Itoa(max) + " index: " + strconv.Itoa(i)
		}
	}

	return minString + maxString
}
