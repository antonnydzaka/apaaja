package main

import "fmt"

func main() {
	var min, max int = 0, 100
	var jawaban bool
	var angka = 64
	for max != min {
		if angka < min {
			fmt.Print(min)
		} else {
			fmt.Print(angka)
		}
		fmt.Scan(&jawaban)
		if jawaban {
			min += angka
		} else {
			max -= angka
		}
		angka /=2
	}
	fmt.Print("hasil : ", min)
}
