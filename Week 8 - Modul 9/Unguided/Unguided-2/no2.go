package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&arr[i])
	}

	// a
	fmt.Println("\nIsi array:")
	fmt.Println(arr)

	// b
	fmt.Println("\nIndeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	// c
	fmt.Println("\n\nIndeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	// d
	var x int
	fmt.Print("\n\nKelipatan indeks x = ")
	fmt.Scan(&x)

	fmt.Println("Elemen indeks kelipatan", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// e
	var hapus int
	fmt.Print("\n\nHapus indeks: ")
	fmt.Scan(&hapus)

	arr = append(arr[:hapus], arr[hapus+1:]...)

	fmt.Println("Array setelah dihapus:")
	fmt.Println(arr)

	// f rata-rata
	total := 0
	for _, v := range arr {
		total += v
	}

	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata =", rata)

	// g standar deviasi
	var jumlah float64
	for _, v := range arr {
		jumlah += math.Pow(float64(v)-rata, 2)
	}

	sd := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi =", sd)

	// h frekuensi
	var cari, frek int
	fmt.Print("Cari frekuensi angka: ")
	fmt.Scan(&cari)

	for _, v := range arr {
		if v == cari {
			frek++
		}
	}

	fmt.Println("Frekuensi =", frek)
}