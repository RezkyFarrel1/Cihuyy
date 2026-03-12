package main

import (
	"fmt"
)

func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		return
	}

	fmt.Printf("%d %d\n", permutasi(a, c), kombinasi(a, c))

	fmt.Printf("%d %d\n", permutasi(b, d), kombinasi(b, d))
}