package main

import (
	"fmt"
	"math"
)

func main() {
	var r int

	
	fmt.Print("Jari-jari = ")
	fmt.Scan(&r)

	
	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(r), 3)
	
	
	luas := 4.0 * math.Pi * math.Pow(float64(r), 2)

	
	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}