package main

import ("fmt")

type User struct {
	Name string
	Mail string
	Age int8
}

func ChangeMail(input *User, changed string) {
	input.Mail = changed
}

func main() {
	var person_a User = User{
		Name: "Ridho",
		Mail: "my@mail.com",
		Age: 17,
	}

	fmt.Println(person_a)
	// Argument should be a pointer!
	ChangeMail(&person_a, "test@mail.com")
	fmt.Println(person_a)


}