package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var daftarPemenang []string 

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			daftarPemenang = append(daftarPemenang, klubA)
		} else if skorA < skorB {
			daftarPemenang = append(daftarPemenang, klubB)
		} else {
			daftarPemenang = append(daftarPemenang, "Draw")
		}
		
		pertandingan++
	}

	for i, pemenang := range daftarPemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, pemenang)
	}
	
	fmt.Println("Pertandingan selesai")
}