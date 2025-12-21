package main
import "fmt"
func main(){
	var apa = [...]string{"apel","jeruk","melon"}
	for i, apakek := range apa{
		fmt.Print(i,apakek)
	}
}