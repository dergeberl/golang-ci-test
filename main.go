package main

import (
	"fmt"
)

var version = "dev"
var date = "none"
var commit = "none"

func main() {
	fmt.Println("Hello World")
	HelloWorld()
	// test go vet
	//var world string
	//go func() {
	//	world = "hello"
	//}()

	//gosec test
	username := "admin"
	password := "CQP7qcDT6DZqh2qoqpvTTz8"
	_ = fmt.Sprintf("%v %v", username, password)
	fmt.Printf("%v %v %v\n", version, date, commit)
}

//HelloWorld returns Hello World
func HelloWorld() {
	fmt.Println("Hello World from func")
}
