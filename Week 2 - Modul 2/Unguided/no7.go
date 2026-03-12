package main

import "fmt"

func main() {
	var b int
	count := 0

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}

	fmt.Println()

	if count == 2 {
		fmt.Println("prima: true")
	} else {
		fmt.Println("prima: false")
	}
}
