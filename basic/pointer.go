package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func changeCountry(address *Address){
	address.Country = "Manchester"
}

func main(){
	address1 := Address{"Jakarta","DKI","Indonesia"}
	address2 := &address1 // with pointer

	address2.City = "Bogor"
	*address2 = Address{"Bekasi","Jawa Barat","Indonesia"}
	address4 := new(Address)
	address4.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address4)

	addree3 := Address{
		City: "London",
		Province: "Blue",
		Country: "",
	}

	changeCountry(&addree3)
	fmt.Println(addree3)
}
