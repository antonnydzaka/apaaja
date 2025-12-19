package main
import "fmt"
func main(){
	var angka,total,banyak float32
	for angka != 9999{
		fmt.Scan(&angka)
		total += angka 
		banyak ++
		if angka == 9999{
			total -= 9999
			banyak -= 1
		}
	}
	fmt.Print(total/banyak)
}