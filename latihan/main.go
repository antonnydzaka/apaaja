package main 
import "fmt"
func main (){
	var avg = ratarata(1,2,3,4,5,6,7,4,3,2,2)
	fmt.Print(avg)
}
func ratarata(vels...int)float32{
	var total int 
	for _,vel:= range vels{
		total += vel
	}
	var avg float32= float32(total)/float32(len(vels))
	return avg
}