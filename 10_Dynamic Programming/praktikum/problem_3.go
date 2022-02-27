package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}

func Frog(jumps []int) int {
	result := []int{}
	result = append(result, 0)
	result = append(result, int(math.Abs(float64(jumps[0]-jumps[1]))))

	for i := 2; i < len(jumps); i++ {
		minSatu := int(math.Abs(float64(jumps[i-1] - jumps[i])))
		minDua := int(math.Abs(float64(jumps[i-2] - jumps[i])))

		result = append(result, func(a, b int) int {
			if a < b {
				return a
			} else {
				return b
			}
		}(minSatu+result[i-1], minDua+result[i-2]))
	}

	return result[len(result)-1]
}
