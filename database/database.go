package database

import ("fmt")
var connection string

func init() {
	connection = "Mysql"
	fmt.Println("Init is executed!")
}

func GetDatabase() string {
	return connection
}