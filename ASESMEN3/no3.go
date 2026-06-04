package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int
	var suara int

	fmt.Println("Masukkan proses input suara :")
	
	fmt.Scan(&suara)
	for suara != -1 {
		idx := posisi(p, n, suara)
		if idx != -1 {
			p[idx].suara++
		} else {
			p[n].nama = suara
			p[n].suara = 1
			n++
		}
		fmt.Scan(&suara)
	}

	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	fmt.Println("\nHasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}