package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

		fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\n--- HASIL OPERASI ARRAY ---")

	fmt.Println("a. Seluruh isi array:", arr)

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("   Elemen dengan indeks kelipatan %d: ", x)
	if x > 0 {
		for i := 0; i < len(arr); i += x {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idxHapus int
	fmt.Print("e. Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	if idxHapus >= 0 && idxHapus < len(arr) {
		arr = append(arr[:idxHapus], arr[idxHapus+1:]...)
		fmt.Println("   Isi array setelah dihapus:", arr)
	} else {
		fmt.Println("   Indeks tidak valid!")
	}

		if len(arr) == 0 {
		fmt.Println("\nArray kosong, tidak dapat menghitung rata-rata, standar deviasi, dan frekuensi.")
		return
	}

	var total float64
	for _, nilai := range arr {
		total += float64(nilai)
	}
	rataRata := total / float64(len(arr))
	fmt.Printf("f. Rata-rata dari array: %.2f\n", rataRata)

	var totalKuadratSelisih float64
	for _, nilai := range arr {
		selisih := float64(nilai) - rataRata
		totalKuadratSelisih += selisih * selisih
	}
	standarDeviasi := math.Sqrt(totalKuadratSelisih / float64(len(arr)))
	fmt.Printf("g. Standar Deviasi: %.2f\n", standarDeviasi)

		var bilanganCari int
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilanganCari)
	
	frekuensi := 0
	for _, nilai := range arr {
		if nilai == bilanganCari {
			frekuensi++
		}
	}
	fmt.Printf("   Frekuensi bilangan %d di dalam array: %d kali\n", bilanganCari, frekuensi)
}