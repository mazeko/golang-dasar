package main

import "fmt"

func main(){
	var names[3] string
	names[0] = "Eko"
	names[1] = "Purnomo"
	names[2] = "Abc"

	fmt.Println(names)

	var values = [3]int{
		90,95,80,
	}

	fmt.Println(values)
}