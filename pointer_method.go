package main

import ("fmt")

type Person struct {
	Name string
}

// Use * to sign as pointer
func (person *Person) ChangePrefixName() {
	person.Name = "Sir " + person.Name
}

func main() {
	person_a := Person{
		Name: "Ridho",
	}
	fmt.Println(person_a)
	person_a.ChangePrefixName()
	fmt.Println(person_a)


}