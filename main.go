package main
import (
	"fmt"
	"math/rand"
)
func main(){
	var n int
	var x, y float32
	var A,B,C,D float32
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x = rand.Float32()
		y = rand.Float32()
		if x < 0.5 && y < 0.5 {
			A++
		} else if x >= 0.5 && y < 0.5 {
			B++
		} else if x >= 0.5 && y >= 0.5 {
			C++
		} else if x < 0.5 && y >= 0.5 {
			D++
		}
	}
	fmt.Printf("Daerah A: %f tetes (%.4f ml)\n", A, float64(A)*0.0001)
	fmt.Printf("Daerah B: %f tetes (%.4f ml)\n", B, float64(B)*0.0001)
	fmt.Printf("Daerah C: %f tetes (%.4f ml)\n", C, float64(C)*0.0001)
	fmt.Printf("Daerah D: %f tetes (%.4f ml)\n", D, float64(D)*0.0001)

}
