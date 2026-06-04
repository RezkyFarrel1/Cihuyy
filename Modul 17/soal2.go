package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan string x: ")
    x, err := reader.ReadString('\n')
    if err != nil {
        fmt.Fprintln(os.Stderr, "Gagal membaca x:", err)
        return
    }
    x = strings.TrimSpace(x)

    fmt.Print("Masukkan nilai n: ")
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Fprintln(os.Stderr, "Gagal membaca n:", err)
        return
    }
    line = strings.TrimSpace(line)
    n, err := strconv.Atoi(line)
    if err != nil || n < 0 {
        fmt.Fprintln(os.Stderr, "n harus berupa bilangan bulat positif")
        return
    }

    count := 0
    firstPos := -1
    for i := 1; i <= n; i++ {
        fmt.Printf("Masukkan string ke-%d: ", i)
        s, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, "Gagal membaca string:", err)
            return
        }
        s = strings.TrimSpace(s)

        if s == x {
            count++
            if firstPos == -1 {
                firstPos = i
            }
        }
    }

    ada := count > 0
    minimalDua := count >= 2

    fmt.Println("\nHasil:")
    fmt.Printf("a. Apakah string x ada? %v\n", boolToYesNo(ada))
    if ada {
        fmt.Printf("b. Posisi pertama string x ditemukan: %d\n", firstPos)
    } else {
        fmt.Println("b. Posisi pertama string x ditemukan: -")
    }
    fmt.Printf("c. Jumlah string x dalam data: %d\n", count)
    fmt.Printf("d. Ada sedikitnya dua string x? %v\n", boolToYesNo(minimalDua))
}

func boolToYesNo(value bool) string {
    if value {
        return "Ya"
    }
    return "Tidak"
}
