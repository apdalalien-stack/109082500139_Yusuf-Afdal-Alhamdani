package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string
	no := 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d : ", no)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		no++
	}

	fmt.Println()

	for i := 0; i < len(hasil); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}