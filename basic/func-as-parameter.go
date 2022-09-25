package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	nameFilterd := filter(name)
	fmt.Println("Hallo ", nameFilterd)
}

func spamFilter(name string) string {
	if name == "Gblk" {
		return "..."
	}

	return name
}

func main(){
	sayHelloWithFilter("Gblk", spamFilter)
}