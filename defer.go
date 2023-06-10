package main

import ("fmt")

func foo() {
	message := recover()
	fmt.Println("Error message: ", message)
	fmt.Println("Executed after other")
}


func bar() {
	defer foo()
	fmt.Println("Waiting...")
	
	
}

func testing(stats bool) {
	defer foo()
	if stats {
		panic("App is stopping")
	}
	
}

func main() {
	testing(true)
	// foo()
}