package main

import "fmt"

func main(){
	var name1 = "Eko"
	var name2 = "Eko"

	var result bool = name1 == name2
	fmt.Println(result)

	var angka1 = 100
	var angka2 = 200
	fmt.Println(angka1 > angka2)
	fmt.Println(angka1 < angka2)
	fmt.Println(angka1 == angka2)
	fmt.Println(angka1 != angka2)
}