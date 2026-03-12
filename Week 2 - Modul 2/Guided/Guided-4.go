package main
import "fmt"

func main() {
	var c float64
	var r, f, k float64

	fmt.Print("Masukkan Celsius: ")
	fmt.Scanln(&c)

	r = c * 4 / 5
	f = (c * 9 / 5) + 32
	k = c + 273

	fmt.Printf("Derajat Reamur: %.0f\n", r)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", f)
	fmt.Printf("Derajat Kelvin: %.0f\n", k)
}