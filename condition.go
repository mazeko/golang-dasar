package main

import "fmt"

func main(){
	var name = "Ekonomi"

	if name == "Eko"{
		fmt.Println(name)
	}

	// short statement
	if length := len(name); length >= 5{
		fmt.Println("Hallo ", name)
	}else{
		fmt.Println("Kurang panjang")
	}
}