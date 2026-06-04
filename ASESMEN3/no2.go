package main

import "fmt"

type Player struct {
	firstName string
	lastName  string
	goals     int
	assists   int
}

type arrPlayer [1001]Player

func SelectionSort(A *arrPlayer, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i+1; j < n; j++ {
			if A[j].goals > A[maxIdx].goals {
				maxIdx = j
			} else if A[j].goals == A[maxIdx].goals {
				if A[j].assists > A[maxIdx].assists {
					maxIdx = j
				}
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
}

func main() {
	var n int
	var players arrPlayer

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&players[i].firstName, &players[i].lastName, &players[i].goals, &players[i].assists)
	}

	SelectionSort(&players, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", players[i].firstName, players[i].lastName, players[i].goals, players[i].assists)
	}
}