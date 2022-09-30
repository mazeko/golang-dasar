package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jalan jalan"
	eko.Age = 25

	fmt.Println(eko)

	joko := Customer {
		Name: "Joko",
		Address: "jakarta",
		Age: 25,
	}

	fmt.Println(joko)
}