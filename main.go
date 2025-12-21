package main
import "fmt"
func main(){
	var buah = []string {"apel","nanas","jeruk"}
	fmt.Print(buah[:])
	var bbuah = buah[:]
	bbuah[0] = "pepaya"
	fmt.Print(bbuah)
}