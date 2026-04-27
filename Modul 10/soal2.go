package main

import (
	"fmt"
)

func main() {
	var beratIkan [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	totalSemuaBerat := 0.0
	beratWadahSementara := 0.0
	jumlahWadah := 0

	for i := 0; i < x; i++ {
		beratWadahSementara += beratIkan[i]
		totalSemuaBerat += beratIkan[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Print(beratWadahSementara, " ")
			jumlahWadah++
			beratWadahSementara = 0
		}
	}
	fmt.Println()

	rataRataWadah := totalSemuaBerat / float64(jumlahWadah)
	fmt.Println(rataRataWadah)
}