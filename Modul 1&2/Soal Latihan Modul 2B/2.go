package main

import (
	"fmt"
)

func main() {
	var pita, bunga string
	var jumlahBunga int

	for i := 1; ; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
		jumlahBunga++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", jumlahBunga)
}