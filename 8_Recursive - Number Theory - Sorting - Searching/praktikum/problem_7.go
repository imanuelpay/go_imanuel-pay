package main

import "fmt"

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}

func playingDomino(card [][]int, deck []int) interface{} {
	index := -1
	for i := 0; i < len(card); i++ {
		for j := 0; j < len(card[i]); j++ {
			if card[i][j] == deck[0] || card[i][j] == deck[1] {
				index = i
				break
			}
		}

		if index > -1 {
			break
		}
	}

	if index > -1 {
		return card[index]
	}

	return "tutup kartu"
}
