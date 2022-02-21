package main

import "fmt"

func main() {
	playWithAsterik(5)
}

func playWithAsterik(n int) {
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			fmt.Printf(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}
