package main
import "fmt"
func main(){
	var target,total,donasi int
	var donatur int 
	fmt.Scan(&target)
	for target > total{
		donatur ++
		fmt.Print("donatur ", donatur," menyumbang sebesar: ")
		fmt.Scan(&donasi)
		total += donasi
		fmt.Println("total donasi",total)
	}
	fmt.Println("target tercapai total donasi: ",total,"dari ",donatur)
}