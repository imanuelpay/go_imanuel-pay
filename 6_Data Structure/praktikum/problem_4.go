package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Output:", munculSekali("1234123"))
	fmt.Println("Output:", munculSekali("76523752"))
	fmt.Println("Output:", munculSekali("12345"))
	fmt.Println("Output:", munculSekali("1122334455"))
	fmt.Println("Output:", munculSekali("0872504"))
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
