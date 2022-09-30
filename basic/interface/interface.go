package main

import "fmt"

type HasName interface{
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func sayHola(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

func main(){
	var eko Person 
	eko.Name = "Eko"

	sayHola(eko)

	cat := Animal {
		Name: "Meow",
	}

	sayHola(cat)
}