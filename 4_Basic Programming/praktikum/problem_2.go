package main

import "fmt"

func main() {
	var studentScore int = 80

	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("Nilai A")
	} else if studentScore >= 65 && studentScore < 80 {
		fmt.Println("Nilai B")
	} else if studentScore >= 50 && studentScore < 65 {
		fmt.Println("Nilai C")
	} else if studentScore >= 35 && studentScore < 50 {
		fmt.Println("Nilai D")
	} else if studentScore < 35 && studentScore >= 0 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
