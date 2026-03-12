package main
import "fmt"

func main() {
	var a, b float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 || a+b > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if a-b >= 9 || b-a >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}