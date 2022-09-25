package main

import "fmt"

func main(){
	counter := 1

	for counter < 3 {
		fmt.Println("Angka ", counter)

		counter++
	}

	// with statement
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	var slice = []string{"Programmer","Zaman","Now"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index ", i,"=", value)
		fmt.Println("Value", value)
	}
}