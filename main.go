package main
import "fmt"
func main(){
	var poin int 
	var hasil bool = false
	var i int 
	var pertandingan int = 0
	for i != 5{
		fmt.Scan(&poin)
		if poin > 500{
			hasil = true
			break
		}
		if poin > 0{
			pertandingan += 1
		}
		i++
	}
	if pertandingan == 5{
		hasil = true
	}
	fmt.Print(hasil)
}