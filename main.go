package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	HelloWorld()
	// test go vet
	var world string
	go func() {
		world = "hello"
	}()
	//gosec test
	username := "admin"
	password := "CQP7qcDT6DZqh2qoqpvTTz8"
	_ = fmt.Sprintf("%v %v", username, password)
}

func HelloWorld(){
	fmt.Println("Hello World from func")
}