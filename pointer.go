package main

import ("fmt")

type Address struct {
	City string
	Province string
	Country string
}

func main() {
	var address1 Address= Address{
		City: "Sukoharjo",
		Province: "Jawa Tengah",
		Country: "Indonesia",
	}

	var address2 *Address= &address1
	var address3 *Address = &address1

	*address2 = Address{"Jogja", "DKI Jogjakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}