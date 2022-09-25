package main

import "fmt"

func logging(){
	fmt.Println("Logg")
}

func runApp(value int){
	defer logging()
	fmt.Println("Running")
	result = 10 / value
	fmt.Println("Value ", result)
	// logging()
}

func main(){
	runApp(0)
}