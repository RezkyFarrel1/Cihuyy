package main

import (
	"fmt"
)

func main() {
	var suaraMasuk int 
	var suaraSah int   
	var rekapSuara [21]int 

	for {
		var suara int
		_, err := fmt.Scan(&suara)
		
		if err != nil {
			break
		}

		if suara == 0 {
			break
		}

		suaraMasuk++

		if suara >= 1 && suara <= 20 {
			suaraSah++
			rekapSuara[suara]++ 
		}
	}

	fmt.Printf("Suara masuk : %d\n", suaraMasuk)
	fmt.Printf("Suara sah : %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if rekapSuara[i] > 0 {
			fmt.Printf("%d : %d\n", i, rekapSuara[i])
		}
	}
}