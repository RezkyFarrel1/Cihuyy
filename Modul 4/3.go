package main

import (
	"fmt"
)

func cetakDeret(n int) {
	fmt.Print(n)

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf(" %d", n)
	}
	
	fmt.Println()
}

func main() {
	var n int
	
	fmt.Scan(&n)

	cetakDeret(n)
}