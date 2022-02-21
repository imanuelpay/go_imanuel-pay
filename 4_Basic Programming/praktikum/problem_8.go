package main

import "fmt"

func main() {
	cetakTabelPerkalian(9)
}

func cetakTabelPerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d\t", j*i)
		}
		fmt.Println()
	}
}
