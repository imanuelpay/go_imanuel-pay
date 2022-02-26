package main

import "fmt"

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)
}

func BinarySearch(array []int, x int) {
	result, kiri, kanan := -1, 0, len(array)-1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if array[tengah] == x {
			result = tengah
		}

		if array[tengah] < x {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println(result)
}
