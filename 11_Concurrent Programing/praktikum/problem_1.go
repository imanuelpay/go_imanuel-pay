package main

import (
	"fmt"
	"time"
)

func main() {
	words := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	wordsChan := make(chan string)
	count := map[string]int{}
	lenWord := 0

	for i := 0; i < len(words); i++ {
		if words[i] == 32 || words[i] == 44 || words[i] == 46 {
			continue
		}

		go func(word string) {
			wordsChan <- word
		}(string(words[i]))

		lenWord++
	}

	go func() {
		for i := 0; i < lenWord; i++ {
			word := <-wordsChan

			value, cek := count[word]
			if cek {
				count[word] = value + 1
			} else {
				count[word] = 1
			}

			fmt.Println(word, ": ", count[word])
		}
	}()

	// defer func() {
	// 	for word, value := range count {
	// 		fmt.Println(word, ": ", value)
	// 	}
	// }()

	<-time.After(1 * time.Second)
}
