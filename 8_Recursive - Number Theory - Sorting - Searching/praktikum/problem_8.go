package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var mapItems = map[string]int{}
	for _, value := range items {
		mapItems[value] = 0
	}

	for _, value := range items {
		data, cek := mapItems[value]
		if cek {
			mapItems[value] = data + 1
		}
	}

	var p = []pair{}
	for key, value := range mapItems {
		p = append(p, pair{name: key, count: value})
	}

	sort.SliceStable(p, func(i, j int) bool {
		return p[i].count < p[j].count
	})

	return p
}
