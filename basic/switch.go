package main

import "fmt"

func main(){
	var name = "Jokowi"

	switch(name){
		case "Eko":
			fmt.Println("Hallo", name)
		case "Joko":
			fmt.Println("Hallo", name)
		default:
			fmt.Println("Gak kenal")
	}
}