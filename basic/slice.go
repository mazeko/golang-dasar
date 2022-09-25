package main

import "fmt"

func main(){
	var months = []string{
		"Januari","Feb","Maret","April","Mei","Juni","Juli","Aug","Sept","Okt","Nov","Des",
	}
	fmt.Println(months)

	var slice1 = months[4:7]

	fmt.Println(slice1)

	months[5] = "Updated"
	fmt.Println(slice1)

	slice1[0] = "Cool"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Updated"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)
}