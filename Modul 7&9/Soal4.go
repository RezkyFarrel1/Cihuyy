package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	var input string
	for {
		fmt.Scan(&input)

		if input == "." {
			break
		}

		for _, r := range input {
			if r == '.' {
				return
			}
			if *n < NMAX {
				t[*n] = r
				*n = *n + 1
			}
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var tBalik tabel = t

	balikanArray(&tBalik, n)

	for i := 0; i < n; i++ {
		if t[i] != tBalik[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks \t\t: ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks \t: ")
	var tabReverse tabel = tab
	balikanArray(&tabReverse, m)
	cetakArray(tabReverse, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom \t? %v\n", isPalindrom)
}