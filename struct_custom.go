package main

import "fmt"

// Custom type definition
// Age is a custom type based on int
type Age int

// User struct with custom type
type User struct {
	Name  string
	Age   Age
	Email string
}

// Function to demonstrate struct and custom type usage
func StructCustomDemo() {
	user := User{
		Name:  "Abhishek",
		Age:   24,
		Email: "abhishek@example.com",
	}
	fmt.Printf("User: %+v\n", user)
	// Using custom type directly
	var myAge Age = 30
	fmt.Println("My Age (custom type):", myAge)
}
