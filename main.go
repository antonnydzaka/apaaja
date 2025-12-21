package main

import "fmt"

func main() {
	var angka = []int{1, 5, 7, 9, 2}
	var penjumlahan int
	for i := 0; i < len(angka); i++ {
		penjumlahan += angka[i]
	}
	fmt.Println(penjumlahan)
	var average int = penjumlahan / len(angka)
	fmt.Println(average)
	var max, min int = angka[0], angka[0]
	for i := range angka {
		if max < angka[i] {
			max = angka[i]
		} else if min > angka[i] {
			min = angka[i]
		}
	}
	fmt.Println(max, min)
}
