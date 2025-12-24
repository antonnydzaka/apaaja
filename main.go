package main

import "fmt"

func main() {
	var input string
	translate := map[string]string{
		"home": "rumah",
		"cat":  "kucing",
		"cap":  "topi",
		"hight":"tinggi",
		"dog":"anjing",
	}
	fmt.Scan(&input)
	var value,ada = translate[input]
	if ada{
		fmt.Print(value)
	}else{
		fmt.Print("tidak ada didalam kamus")
	}

}
