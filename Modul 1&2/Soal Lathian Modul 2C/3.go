package main

import (
	"fmt"
)

func main() {
	var b int
	var jumlahFaktor int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	
	
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i) 
			jumlahFaktor++      
		}
	}
	
	fmt.Println() 

	
	isPrima := (jumlahFaktor == 2)
	fmt.Printf("Prima: %t\n", isPrima)
}