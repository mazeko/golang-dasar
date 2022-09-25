package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList){
	if blacklist(name){
		fmt.Println("Blocked")
	}else{
		fmt.Println("welcome")
	}
}

func main(){
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Eko", blacklist)
}