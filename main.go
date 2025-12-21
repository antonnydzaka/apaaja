package main
import "fmt"
func main(){
	var buah = []string {"apel","nanas","jeruk","melon","semangka"}
	fmt.Println(buah[:])
	var bbuah = buah[0:3]
	bbuah[0] = "pepaya"
	var buuah = buah 
	buuah = append(buuah,"pir")
	fmt.Println(bbuah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	fmt.Println(len(bbuah))
	fmt.Println(cap(bbuah))
}