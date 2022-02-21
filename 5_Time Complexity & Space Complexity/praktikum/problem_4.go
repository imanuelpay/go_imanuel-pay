package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)

	fmt.Println("Output:", munculSekali(s))
}

func munculSekali(s string) []int {
	var result []int
	for i := 0; i < len(s); i++ {
		var count = 0

		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count += 1
			}
		}

		if count < 2 {
			k, err := strconv.Atoi(string(s[i]))
			if err != nil {
				panic(err)
			}
			result = append(result, k)
		}
	}

	return result
}
