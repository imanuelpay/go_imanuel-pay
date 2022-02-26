package main

import (
	"fmt"
	"sort"
)

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}

func DragonOfLoowater(dragonHead, knightHeight []int) {
	dragonPower := 0
	minimumKnightHeight := []int{}
	for _, value := range dragonHead {
		dragonPower += value
	}

	sort.Ints(knightHeight)
	if len(knightHeight) <= 1 {
		minimumKnightHeight = append(minimumKnightHeight, knightHeight[0])
	} else {
		for i := 0; i < len(knightHeight); i++ {
			for j := 0; j < len(knightHeight); j++ {
				if knightHeight[i]+knightHeight[j] > dragonPower && i != j {
					minimumKnightHeight = append(minimumKnightHeight, (knightHeight[i] + knightHeight[j]))
				}
			}
		}
	}

	sort.Ints(minimumKnightHeight)
	if len(minimumKnightHeight) < 1 || minimumKnightHeight[0] <= dragonPower {
		fmt.Println("Knight fall")
	} else {
		fmt.Println(minimumKnightHeight[0])
	}
}
