package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var karakter string

	
	fmt.Scan(&a, &b, &c, &d, &e)
	
	
	fmt.Scan(&karakter)

	
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	for i := 0; i < len(karakter); i++ {
		fmt.Printf("%c", karakter[i]+1) 
	}
	fmt.Println()
}