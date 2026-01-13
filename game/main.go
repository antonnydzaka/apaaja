package main
import "fmt"
func main(){
	var min,max int = 0,100
	var jawab bool
	fmt.Println("--selamaat datang--")
	fmt.Println("--pikirkan angka 0 sampai 100--")
	fmt.Println("--press enter--")
	fmt.Scanln( )
	for min != max{
		var x int = (min+max)/2
		fmt.Println("apakah angka lebih besar dari ",x)
		fmt.Println("[true/false]")
		fmt.Scan(&jawab)
		if jawab{
			min = x + 1
		}else{
			max = x
		}
	}
	fmt.Println("apakah memikirkan angka ",min,"[true/flase]")
	var terakhir bool
	fmt.Scan(&terakhir)
	if terakhir{
		fmt.Println("horee")
	}else{
		fmt.Println("yahh salah")
	}
}