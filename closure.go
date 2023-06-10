package main

func main() {
	counter := 0
	foo := func(a int) int {
		counter = 5
		return counter
	}

	fmt.Println(foo)

}