package main

import "fmt"

func main(){
	var nilai32 int32 = 12345
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	
	var name = "Eko"
	// var estring = string(name[1])

	fmt.Println(name)
	fmt.Println(string(name[1]))

}