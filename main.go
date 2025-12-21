package main

import "fmt"

func main(){
	var x,y int
	fmt.Scan(&x,&y)
	var fpb []int
	var kpk []int
	for i := 1; i <= x || i <= y; i++ {
		if x%i==0 && y%i==0{
			fpb = append(fpb, i)
			x/=i
			y/=i
		}
	}
	for i := 1; i <= x || i <= y; i++ {
		if x%i==0 || y%i==0{
			kpk = append(kpk,i)
			x/=i
			y/=i
		}
	}
	fmt.Println(fpb," ")
	fmt.Println(kpk," ")
	fmt.Println(len(fpb))
	fmt.Println(cap(fpb))
}