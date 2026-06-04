package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan banyak topping: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Gagal membaca input:", err)
		return
	}

	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil || n < 0 {
		fmt.Fprintln(os.Stderr, "Banyak topping harus berupa bilangan bulat non-negatif")
		return
	}

	rand.Seed(time.Now().UnixNano())
	var hitInside, hitOutside int
	centerX, centerY := 0.5, 0.5
	radius := 0.5
	radius2 := radius * radius

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		dx := x - centerX
		dy := y - centerY
		if dx*dx+dy*dy <= radius2 {
			hitInside++
		} else {
			hitOutside++
		}
	}

	fmt.Printf("Topping pada pizza: %d\n", hitInside)
	fmt.Printf("Topping pada wadah: %d\n", hitOutside)
	fmt.Printf("Perbandingan pizza: %.8f\n", float64(hitInside)/float64(n))
	fmt.Printf("Perbandingan wadah: %.8f\n", float64(hitOutside)/float64(n))
}
