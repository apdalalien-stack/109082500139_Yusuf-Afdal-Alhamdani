package main

import "fmt"

func main() {
	var arr []int
	var nomer int

	for {
		fmt.Scan(&nomer)
		if nomer < 0 {
			break
		}
		arr = append(arr, nomer)
	}
	Insertion(arr)

	for i := 0; i < len(arr); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	if len(arr) < 2 {
		fmt.Println("Data berjarak 0") // Jika data kurang dari 2 elemen
	} else {
		diff := arr[1] - arr[0]
		isTetap := true

		for i := 2; i < len(arr); i++ {
			if arr[i]-arr[i-1] != diff {
				isTetap = false
				break
			}
		}

		if isTetap {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
