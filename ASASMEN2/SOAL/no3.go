package main

import "fmt"

const nProv = 10

type NamaProv [1 + nProv]string
type PopProv [1 + nProv]int
type TumbuhProv [1 + nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 1; i <= nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 1

	for i := 2; i <= nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}

	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var prediksi int

	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i := 1; i <= nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi = int((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 1; i <= nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string
	var idx int

	fmt.Println("Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi")

	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cari)

	idx = ProvinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idx])

	idx = IndeksProvinsi(prov, cari)
	fmt.Println("\nData provinsi yang dicari :", prov[idx])

	fmt.Println()
	Prediksi(prov, pop, tumbuh)
}