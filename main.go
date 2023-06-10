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

type Orang struct {
	Nama string
}

type Hewan struct {
	Nama string
}

// Interface
type HasName interface {
	GetName() string
}


func SayHi(hasName HasName) {
	fmt.Println("Hi ", hasName.GetName())
}

func (orang Orang) GetName() string {
	return orang.Nama
}

// Struct method
func (student Student) greet(teacher string) {
	fmt.Printf("Hello, %v. Your age is %v. This is your teacher, %v \n", student.Name, student.Age, teacher)
} 

func (animal Hewan) GetName() string {
	return animal.Nama
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

	fmt.Println(border)

	seseorang := Orang{
		Nama : "Budi",
	}

	SayHi(seseorang)

	neko := Hewan{
		Nama : "Kucing",
	}

	SayHi(neko)

}