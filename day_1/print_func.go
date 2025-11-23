package main

import "fmt"

func printMain() {
	fmt.Print("Hello", "World!")   //without new line and space
	fmt.Println("Hello", "World!") //with new line and space
	var (
		name, age = "John", 25
	)
	fmt.Printf("Name: %s, Age: %d", name, age) //with new line and space and format

}
