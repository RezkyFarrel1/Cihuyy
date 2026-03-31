package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Program Deret Fibonacci (Rekursif)")
	fmt.Println("----------------------------------")
	fmt.Println("n\tSn")
	fmt.Println("----------------------------------")

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t%d\n", i, fibonacci(i))
	}
}