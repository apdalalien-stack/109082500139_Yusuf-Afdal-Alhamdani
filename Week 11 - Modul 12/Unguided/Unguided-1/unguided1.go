package main

import "fmt"

type suara struct {
	indxrt      int
	indxpilkart int
}

func main() {
	var rt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var pilkart []int
	var hasil []suara
	var input int = -1
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		} else {
			if input > 0 || input <= 20 {
				pilkart = append(pilkart, input)
			}
		}
	}
	for i := 0; i < len(rt); i++ {
		x := validate(pilkart, rt[i])
		if x != 0 {
			hasil = append(hasil, suara{indxrt: rt[i], indxpilkart: x})
		}
	}
	fmt.Println("suara masuk:", len(pilkart))
	j :=0
	for _,vel:= range hasil{
		j += vel.indxpilkart
	}
	fmt.Println("suara sah:",j)
	print(hasil)
}

func validate(pilkart []int, rt int) int {
	var i int
	for _, vel := range pilkart {
		if vel == rt {
			i++
		}
	}
	return i
}

func print(hasil []suara) {
	for _, vel := range hasil {
		fmt.Println(vel.indxrt, ":", vel.indxpilkart)
	}
}
