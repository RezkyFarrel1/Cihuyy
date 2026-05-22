package main

import (
    "bufio"
    "fmt"
    "os"
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
        if x < 0 {
            break
        }
        data = append(data, x)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading input:", err)
        os.Exit(1)
    }

    insertionSort(data)
    printData(data)
}

func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

func printData(data []int) {
    for i, v := range data {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println()

    if len(data) < 2 {
        fmt.Println("Data berjarak tidak tetap")
        return
    }

    diff := data[1] - data[0]
    konstant := true
    for i := 2; i < len(data); i++ {
        if data[i]-data[i-1] != diff {
            konstant = false
            break
        }
    }

    if konstant {
        fmt.Printf("Data berjarak %d\n", diff)
    } else {
        fmt.Println("Data berjarak tidak tetap")
    }
}
