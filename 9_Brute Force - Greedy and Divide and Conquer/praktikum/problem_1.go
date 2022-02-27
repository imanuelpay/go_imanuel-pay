package main

import "fmt"

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}

func SimpleEquations(a, b, c int) {

	// Mencari faktor bilangan dari B
	factor := []int{}
	for i := 1; i <= b; i++ {
		if (b % i) == 0 {
			factor = append(factor, i)
		}
	}

	// Menghitung apakah penjumlahan 3 faktor bilangan B, apakah sama dengan A
	// Jika sama, masukkan 3 faktor bilangan tersebut ke Slice cekJumlah
	cekJumlah := []int{}
	for i := 0; i < len(factor)-2; i++ {
		if (factor[i] + factor[i+1] + factor[i+2]) == a {
			cekJumlah = append(cekJumlah, factor[i], factor[i+1], factor[i+2])
		}
	}

	C := 0
	for _, data := range cekJumlah {
		C += data * data
	}

	// Jika Slice cekJumlah tidak kosong, maka cetak slice cekJumlah
	if len(cekJumlah) < 1 || C != c {
		fmt.Println("No solution.")
	} else {
		fmt.Println(cekJumlah)
	}
}
