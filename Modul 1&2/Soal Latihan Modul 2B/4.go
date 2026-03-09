package main

import (
	"fmt"
	"math"
)

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	akar2 := 1.0

	for k := 0; k <= K; k++ {
	
		pembilang := math.Pow(float64(4*k+2), 2)
		
		
		penyebut := float64(4*k+1) * float64(4*k+3)

	
		akar2 *= (pembilang / penyebut)
	}


	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}