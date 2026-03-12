package main
import "fmt"

func main() {
	var bunga string
	var pita string
	count := 0
	i := 1

	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		count++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}