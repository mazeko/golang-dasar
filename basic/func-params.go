package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Println("hallo ",firstName,lastName)
}

func main(){
	sayHello("Eko","Purnomo")
}