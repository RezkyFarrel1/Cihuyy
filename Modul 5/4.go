package main

import "fmt"

func cetakBarisan(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
		return
	}
	
	fmt.Printf("%d ", n)
	cetakBarisan(n - 1)
	fmt.Printf("%d ", n)
}

func main() {
	var n int
	
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	
	fmt.Print("Keluaran: ")
	cetakBarisan(n)
	fmt.Println()
}