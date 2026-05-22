package main

import (
    "fmt"
)

func selectionSortAsc(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func selectionSortDesc(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] > arr[maxIdx] {
                maxIdx = j
            }
        }
        arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
    }
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    for i := 0; i < n; i++ {
        var m int
        if _, err := fmt.Scan(&m); err != nil {
            return
        }

        rumah := make([]int, m)
        for j := 0; j < m; j++ {
            fmt.Scan(&rumah[j])
        }

        var odd, even []int
        for _, x := range rumah {
            if x%2 != 0 {
                odd = append(odd, x)
            } else {
                even = append(even, x)
            }
        }

        selectionSortAsc(odd)
        selectionSortDesc(even)

        first := true
        for _, x := range odd {
            if !first {
                fmt.Print(" ")
            }
            fmt.Print(x)
            first = false
        }
        for _, x := range even {
            if !first {
                fmt.Print(" ")
            }
            fmt.Print(x)
            first = false
        }
        fmt.Println()
    }
}
