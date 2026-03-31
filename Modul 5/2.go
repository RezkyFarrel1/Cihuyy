package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	
	cetakBintang(n - 1)
	
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	
	fmt.Println("Keluaran:")
	cetakBintang(n)
}