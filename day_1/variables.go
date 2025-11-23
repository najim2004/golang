package main

import (
	"fmt"
)

func variablesMain() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	multipleVariableDeclaration()
	unSpecificVariableDeclaration()
	groupVariableDeclaration()
}

func multipleVariableDeclaration() {

	var a, b, c, d = 1, 2, 3, 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func unSpecificVariableDeclaration() {
	var a, b, c = 1, "John", true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func groupVariableDeclaration() {
	var (
		student1 = "John"
		age      = 25
	)
	fmt.Println(student1)
	fmt.Println(age)
}
