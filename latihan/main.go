package main
import "fmt"
func main(){
	var luas, keliling int = kotak(2,5)
	fmt.Print(luas,keliling)
}
func kotak(panjang int , lebar int )(int,int){
	luas := panjang * lebar
	keliling:= 2*(panjang + lebar)
	return luas , keliling
}
