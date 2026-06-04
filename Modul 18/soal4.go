package main

import "fmt"

type CharMachine struct {
	input []rune
	pos   int
}

func (m *CharMachine) start(s string) {
	m.input = append([]rune(nil), []rune(s)...)
	m.pos = 0
}

func (m *CharMachine) maju() {
	if !m.eop() {
		m.pos++
	}
}

func (m *CharMachine) eop() bool {
	return m.pos >= len(m.input) || m.input[m.pos] == '.'
}

func (m *CharMachine) cc() rune {
	if m.eop() {
		return 0
	}
	return m.input[m.pos]
}

func main() {
	m := CharMachine{}
	s := "ALFAKSALE.BETA"
	m.start(s)

	count := 0
	aCount := 0
	leCount := 0
	prev := rune(0)

	for !m.eop() {
		ch := m.cc()
		count++
		if ch == 'A' {
			aCount++
		}
		if prev == 'L' && ch == 'E' {
			leCount++
		}
		prev = ch
		m.maju()
	}

	fmt.Printf("Total karakter dibaca: %d\n", count)
	fmt.Printf("Jumlah huruf A: %d\n", aCount)
	fmt.Printf("Frekuensi LE: %d\n", leCount)
}
