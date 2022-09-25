package main

import "fmt"

func main(){
	type noKtp string

	var noKtpEko noKtp = "123456"
	fmt.Println(noKtpEko)
}