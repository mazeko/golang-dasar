package main

import "fmt"
import "golang-dasar/basic/database"

func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}