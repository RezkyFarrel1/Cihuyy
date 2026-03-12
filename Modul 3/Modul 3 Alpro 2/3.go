package main

import "fmt"

func diDalamLingkaran(x, y, cx, cy, r int64) bool {
	dx := x - cx
	dy := y - cy
	return dx*dx+dy*dy <= r*r
}

func main() {
	var cx1, cy1, r1 int64
	var cx2, cy2, r2 int64
	var x, y int64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := diDalamLingkaran(x, y, cx1, cy1, r1)
	dalam2 := diDalamLingkaran(x, y, cx2, cy2, r2)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}