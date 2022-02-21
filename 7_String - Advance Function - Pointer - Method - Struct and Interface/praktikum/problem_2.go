package main

import "fmt"

func main() {
	fmt.Println(ceaser(3, "abc"))
	fmt.Println(ceaser(2, "alta"))
	fmt.Println(ceaser(10, "alterraacademy"))
	fmt.Println(ceaser(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(ceaser(1000, "abcdefghijklmnopqrstuvwxyz"))
}

func ceaser(offset int, input string) string {
	offset = offset % 26
	var output string

	for i := 0; i < len(input); i++ {
		x := input[i] + byte(offset)
		if input[i]+byte(offset) > 122 {
			output = output + string(x-26)
		} else {
			output = output + string(x)
		}
	}

	return output
}
