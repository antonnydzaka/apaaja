package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var toping,topingpadapizza int
	fmt.Scan(&toping)
	var xc, yc, r float32 = 0.5, 0.5, 0.5
	for i := 0; i < toping; i++ {
		var x, y float32 = rand.Float32(),rand.Float32()
		dx := x - xc
		dy := y - yc
		if (dx*dx + dy*dy) <= r*r{
			topingpadapizza++
		}
	}
	var pi float32 = 4*float32(topingpadapizza)/float32(toping)
	fmt.Println("toping pada pizza: ",topingpadapizza)
	fmt.Println("phi :",pi)
}