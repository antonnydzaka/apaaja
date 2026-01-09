package main 
import "fmt"
func main(){
	var min,max  int = 0,100
	var jawab bool
	fmt.Println("tebak angka, pikir kan angka dari 0 sampai 100")
	for min != max {
		var mid int = (min+max)/2
		fmt.Println("apakah nilai  lebih besar dari ",mid,"[true/false]")
		fmt.Scan(&jawab)
		if jawab{
			min = mid + 1
		}else{
			max = mid 
		}
	}
	fmt.Println("apakah ini angka yang anda pikirkan :",min)
}