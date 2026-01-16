package main

import "fmt"

func main(){
	var angka =[]int{1,2,3,4,5,6,7,8,9,10}
	afterfilter := filter(angka)
	fmt.Print(afterfilter)
}
func filter(angka []int)[]int{
	var hasil = []int{}
	for _,val  := range angka{
		if val % 2 == 0{
			hasil = append(hasil,val)
		}
	}
	return hasil
}