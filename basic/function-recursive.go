package main

import "fmt"

// With loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// With recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}

	return value * factorialRecursive(value - 1)
}

func main(){
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}