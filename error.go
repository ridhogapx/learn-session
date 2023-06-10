package main

import (
	"fmt"
	"errors"
)

func dividing(num int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider zero is not allowerd")
	} else { 
		result := num / divider
		return result, nil
	}
}

func main() {
	result, err := dividing(50, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
