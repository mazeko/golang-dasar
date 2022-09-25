package main

import "fmt"

func main(){
	var nilaiAkhir = 80
	var nilaiUjian = 75

	var lulusAkhir = nilaiAkhir >= 80
	var lulusUjian = nilaiUjian >= 82

	var lulus = lulusAkhir && lulusUjian
	fmt.Println(lulus)
}