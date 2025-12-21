package main

import "fmt"
func main(){
	type apa string
	var angka string
	fmt.Scan(&angka)
	var digit = angka[0]
	var sdigit = string(digit)
	fmt.Print(sdigit)
}