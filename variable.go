package main

import "fmt"

// Single variable declaration
var Greeting string = "Welcome to Variable in Go"
var Welcome string = "Welcome to Variable in Go 1"

// Multiple variable declarations
var Age int = 24
var City string = "New York"

// Grouped variable declaration
var (
	Country  = "USA"
	Language = "Go"
)

// Function that prints a variable
func variable() {
	var name string = "Abhishek"
	fmt.Println(name)
	email := "mohit2096@gmail.com"
	fmt.Println(email)
	// Assign multiple values in one line
	x, y := 5, "numbers"
	fmt.Println(x, y)
}

// Function that returns a summary strings
func VariableInfo() string {
	return "Name: Abhishek, Age: " + fmt.Sprint(Age) + ", City: " + City + ", Country: " + Country + ", Language: " + Language
}

func CityInformation() string {
	return "City: " + City + ", Country: " + Country
}

// Declare multiple variables
var a, b, c int

// Constants (const)
const Pi = 3.14
const e = 2.71828
const Zero = 0
const Name = "GoLang"

// Block Declaration (Grouped)
