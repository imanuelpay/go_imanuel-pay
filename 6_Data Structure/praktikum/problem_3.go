package main

import "fmt"

func main() {
	fmt.Println("Output:", ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println("Output:", ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println("Output:", ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println("Output:", ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println("Output:", ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println("Output:", ArrayMerge([]string{}, []string{}))
}

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayA = append(arrayA, arrayB...)

	var resultMap = map[string]string{}
	for _, element := range arrayA {
		resultMap[element] = element
	}

	var result []string
	for element := range resultMap {
		result = append(result, element)
	}

	return result
}
