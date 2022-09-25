package main

import "fmt"

func main(){
	var person = map[string]string{
		"name":"Eko",
		"address":"Jakarta",
	}

	fmt.Println(person)

	var book = make(map[string]string)
	book["title"]  = "Go-Lang"
	book["author"] = "Eko" 
	book["release"]= "2022"

	fmt.Println(book)
}