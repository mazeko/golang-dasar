package main

import "fmt"

func main(){
	for i:= 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("================")

	for y := 0; y < 10; y++ {
		if y % 3 == 0 {
			continue
		}
		
		fmt.Println("Perulangan ke ", y)
	}
}