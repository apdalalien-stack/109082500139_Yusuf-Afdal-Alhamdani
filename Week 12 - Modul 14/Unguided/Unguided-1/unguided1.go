package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	input(n)
}

func input(n int){
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}
		selectionSort(arr)
	}

}

func selectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min]{
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
	read(data)
}

func read(data []int){
	for _,vel := range data{
		fmt.Print(vel," ")
	}
	fmt.Println()
}