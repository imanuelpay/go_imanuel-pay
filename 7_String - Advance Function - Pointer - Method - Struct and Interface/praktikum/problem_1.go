package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}

func Compare(a, b string) string {
	cekA := strings.Contains(a, b)
	cekB := strings.Contains(b, a)

	if cekA {
		return b
	} else if cekB {
		return a
	} else {
		return ""
	}
}
