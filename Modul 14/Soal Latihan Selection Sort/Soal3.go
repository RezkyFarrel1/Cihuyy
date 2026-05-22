package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    var data []int

    for scanner.Scan() {
        token := scanner.Text()
        x, err := strconv.Atoi(token)
        if err != nil {
            continue
        }

        switch x {
        case -5313:
            return
        case 0:
            printMedian(data)
        default:
            data = append(data, x)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading input:", err)
        os.Exit(1)
    }
}

func printMedian(data []int) {
    if len(data) == 0 {
        fmt.Println(0)
        return
    }

    sorted := append([]int(nil), data...)
    sort.Ints(sorted)

    n := len(sorted)
    if n%2 == 1 {
        fmt.Println(sorted[n/2])
        return
    }

    fmt.Println((sorted[n/2-1] + sorted[n/2]) / 2)
}
