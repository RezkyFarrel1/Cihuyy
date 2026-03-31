package main

import "fmt"

func cetakGanjil(n int, i int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}

	cetakGanjil(n, i+1)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cetakGanjil(n, 1)
	fmt.Println()
}