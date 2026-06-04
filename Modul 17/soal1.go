	package main

	import (
		"fmt"
	)

	func main() {
		var x float64
		var sum float64
		var count int

		fmt.Println("Masukkan bilangan real, akhiri dengan 9999:")
		for {
			_, err := fmt.Scan(&x)
			if err != nil {
				fmt.Println("Input tidak valid")
				return
			}
			if x == 9999 {
				break
			}
			sum += x
			count++
		}

		if count == 0 {
			fmt.Println("Tidak ada bilangan untuk dihitung reratanya.")
			return
		}

		avg := sum / float64(count)
		fmt.Printf("Rerata: %.2f\n", avg)
	}
