package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	nim   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(A *arrayMahasiswa, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&A[i].nim, &A[i].nama, &A[i].nilai)
	}
}

func cariNilaiPertama(A arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if A[i].nim == nim {
			return A[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(A arrayMahasiswa, n int, nim string) int {
	max := -1

	for i := 0; i < n; i++ {
		if A[i].nim == nim {
			if A[i].nilai > max {
				max = A[i].nilai
			}
		}
	}

	return max
}

func main() {
	var A arrayMahasiswa
	var n int
	var nimCari string
	var nilaiPertama, nilaiTerbesar int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	inputData(&A, n)

	fmt.Print("\nMasukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama = cariNilaiPertama(A, n, nimCari)
	nilaiTerbesar = cariNilaiTerbesar(A, n, nimCari)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
}