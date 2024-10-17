package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(x, y, cx, cy) <= float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	fmt.Print("Masukkan pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	diLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
