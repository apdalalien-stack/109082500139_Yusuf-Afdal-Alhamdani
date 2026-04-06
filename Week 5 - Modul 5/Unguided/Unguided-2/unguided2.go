package main

import "fmt"

func angka(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")

	angka(n - 1) 

	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	angka(n)
}