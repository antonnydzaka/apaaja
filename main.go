package main
import "fmt"
func main(){
	var warna1,warna2,warna3,warna4 string
	var benar int
	var hasil bool
	for i := 1; i <= 5; i++ {
		fmt.Printf("percobaan ke %v :",i)
		fmt.Scan(&warna1,&warna2,&warna3,&warna4)
		if warna1=="merah"&&warna2=="kuning"&&warna3=="hijau"&&warna4=="ungu"{
			benar++
		}
	}
	if benar == 5 {
		hasil = true
	}
	fmt.Print(hasil)
}