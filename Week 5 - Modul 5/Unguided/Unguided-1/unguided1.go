package main

import "fmt"

func cetak(n int) {
	if n == 0 {
		return
	}

	cetak(n - 1) // rekursi dulu

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	cetak(n)
}