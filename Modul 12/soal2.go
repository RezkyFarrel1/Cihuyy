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

	var ketua, wakil int
	var suaraKetua, suaraWakil int = -1, -1

	for i := 1; i <= 20; i++ {
		if rekapSuara[i] > suaraKetua {
			suaraWakil = suaraKetua
			wakil = ketua

			suaraKetua = rekapSuara[i]
			ketua = i
		} else if rekapSuara[i] > suaraWakil {
			suaraWakil = rekapSuara[i]
			wakil = i
		}
	}

	fmt.Printf("Suara masuk : %d\n", suaraMasuk)
	fmt.Printf("Suara sah : %d\n", suaraSah)
	fmt.Printf("Ketua RT : %d\n", ketua)
	fmt.Printf("Wakil ketua : %d\n", wakil)
}