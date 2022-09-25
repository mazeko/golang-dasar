package main

import "fmt"

func getHello(name string) string {
	return "Hallo "+ name
}

func main(){
	fmt.Println(getHello("Eko"))
}