package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan N suku pertama: ")
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Fprintln(os.Stderr, "Gagal membaca input:", err)
        return
    }
    line = strings.TrimSpace(line)
    n, err := strconv.Atoi(line)
    if err != nil || n <= 0 {
        fmt.Fprintln(os.Stderr, "N harus berupa bilangan bulat positif")
        return
    }

    sN := leibnizSum(n)
    fmt.Printf("Hasil PI: %.10f\n", sN)

    tol := 0.00001
    current := sN
    sign := signForTerm(n + 1)
    var next float64
    i := n
    for {
        i++
        term := sign / float64(2*i-1) * 4
        next = current + term
        if math.Abs(next-current) <= tol {
            break
        }
        current = next
        sign = -sign
    }

    fmt.Printf("Hasil PI: %.10f\n", current)
    fmt.Printf("Hasil PI: %.10f\n", next)
    fmt.Printf("Pada i ke: %d\n", i)
}

func leibnizSum(n int) float64 {
    sum := 0.0
    sign := 1.0
    for i := 1; i <= n; i++ {
        sum += sign / float64(2*i-1)
        sign = -sign
    }
    return sum * 4
}

func signForTerm(k int) float64 {
    if k%2 == 1 {
        return 1.0
    }
    return -1.0
}
