package main

import "fmt"

func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Error message: ", message)
	}
	fmt.Println("Finish")
}

func runApps(error bool){
	defer endApp()
	if error {
		panic("Error")
	}

	fmt.Println("Success")
}

func main(){
	runApps(true)
}