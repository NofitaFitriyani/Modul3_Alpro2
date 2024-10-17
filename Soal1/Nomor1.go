package main

import "fmt"

// Fungsi untuk menghitung faktorial
func mencariFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := n; i >= 1; i-- {
		*hasil *= i
	}
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutation(n, r int) int {
	// P(n, r) = n! / (n - r)!
	var pembilang, penyebut int
	mencariFaktorial(n, &pembilang)
	mencariFaktorial(n-r, &penyebut)
	return pembilang / penyebut
}

// Fungsi untuk menghitung kombinasi C(n, r)
func combination(n, r int) int {
	// C(n, r) = n! / (r! * (n - r)!)
	var pembilang, penyebut1, penyebut2 int
	mencariFaktorial(n, &pembilang)
	mencariFaktorial(r, &penyebut1)
	mencariFaktorial(n-r, &penyebut2)
	return pembilang / (penyebut1 * penyebut2)
}

func main() {
	var a, b, c, d int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Baris pertama: Permutasi dan kombinasi dari a terhadap c
	fmt.Printf("P(%d, %d) = %d\n", a, c, permutation(a, c))
	fmt.Printf("C(%d, %d) = %d\n", a, c, combination(a, c))

	// Baris kedua: Permutasi dan kombinasi dari b terhadap d
	fmt.Printf("P(%d, %d) = %d\n", b, d, permutation(b, d))
	fmt.Printf("C(%d, %d) = %d\n", b, d, combination(b, d))
}
