package main

import (
	"fmt"
	"learning_go/helper"
)

type Student struct {
	Name string
	Age int
	Address string
	IsStudent bool
}

func main() {
	var border string = "--------------------------"

	result := helper.Sum(15,20)
	fmt.Println(result)
	fmt.Println(border)

	var foo Student 
	foo.Name = "Ridho"
	foo.Age = 17
	foo.Address = "Sukoharjo"
	foo.IsStudent = true

	fmt.Println(foo.Name)
	fmt.Println(border)

	bar := Student{
		Name : "Joko",
		Age : 28,
		Address: "Sukoharjo",
		IsStudent: false,
	}

	fmt.Println(bar)

}