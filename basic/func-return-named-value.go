package main

import "fmt"

func getFullName()(firstName, middleName, lastName string){
	firstName  = "Programmer"
	middleName = "Zaman"
	lastName   = "Now"

	return
}

func main(){
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}