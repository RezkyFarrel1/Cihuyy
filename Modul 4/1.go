package main

import (
	"fmt"
)

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	
	factorial(n, &nFact)
	factorial(n-r, &nrFact) 
	
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nrFact int
	
	factorial(n, &nFact)     
	factorial(r, &rFact)     
	factorial(n-r, &nrFact)  
	
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Printf("%d %d\n", p1, c1)
	fmt.Printf("%d %d\n", p2, c2)
}