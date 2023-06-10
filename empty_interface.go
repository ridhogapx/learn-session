package main

import ("fmt")

func SayHi() interface{} {
	return 69
}

func main() {
	var result interface{} = SayHi()
	// var resultStr string = result.(string)
	// fmt.Println(resultStr)

	switch val := result.(type) {
	case string:
		fmt.Println("Value", val, "is string")
	case int:
		fmt.Println("Value", val, "is integer")
	default:
		fmt.Println("Value", val, "is unknown")
	}
	
	
}