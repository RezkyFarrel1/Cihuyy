package main

import (
	"fmt"
	"math"
)

func main() {
	var b1, b2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&b1, &b2)

		if b1 < 0 || b2 < 0 || (b1+b2) > 150 {
			break
		}

		oleng := math.Abs(b1-b2) >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
	}

	fmt.Println("Proses selesai.")
}