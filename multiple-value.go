package main

import "fmt"

func getfullName()(string, string){
	return "Eko","Purnomo"
}

func main(){
	firstName , lastName := getfullName()

	fmt.Println(firstName, lastName)
}