package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run soal3.go <jumlah_tetesan>")
        return
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil || n < 0 {
        fmt.Println("Jumlah tetesan harus berupa bilangan bulat non-negatif")
        return
    }

    rand.Seed(time.Now().UnixNano())

    var countA, countB, countC, countD int
    for i := 0; i < n; i++ {
        x := rand.Float64()
        y := rand.Float64()
        if x < 0.5 {
            if y < 0.5 {
                countA++
            } else {
                countD++
            }
        } else {
            if y < 0.5 {
                countB++
            } else {
                countC++
            }
        }
    }

    mmA := float64(countA) * 0.0001
    mmB := float64(countB) * 0.0001
    mmC := float64(countC) * 0.0001
    mmD := float64(countD) * 0.0001

    fmt.Printf("Curah hujan daerah A: %.4f millimeter\n", mmA)
    fmt.Printf("Curah hujan daerah B: %.4f millimeter\n", mmB)
    fmt.Printf("Curah hujan daerah C: %.4f millimeter\n", mmC)
    fmt.Printf("Curah hujan daerah D: %.4f millimeter\n", mmD)
}
