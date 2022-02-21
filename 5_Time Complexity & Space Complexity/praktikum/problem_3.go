package main

import "fmt"

func main() {
	fmt.Println("Output:", arrayMerge([]string{"kazuya", "jin", "lee"}, []string{"kazuya", "feng"}))
	fmt.Println("Output:", arrayMerge([]string{"lee", "jin"}, []string{"kazuya", "panda"}))
}

func arrayMerge(a, b []string) []string {
	a = append(a, b...)

	var resultMap = map[string]string{}
	for _, element := range a {
		resultMap[element] = element
	}

	var result []string
	for element := range resultMap {
		result = append(result, element)
	}

	return result
}
