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

// Struct method
func (student Student) greet(teacher string) {
	fmt.Printf("Hello, %v. Your age is %v. This is your teacher, %v \n", student.Name, student.Age, teacher)
} 

func main() {
	var border string = "---------------------------------------------------"

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

	foo.greet(bar.Name)
	fmt.Println(bar)

}