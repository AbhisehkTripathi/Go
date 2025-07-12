package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

var FunctionGreeting string = "Welcome to Function in Go"

// Basic Function (No Params, No Return)
func function() {
	fmt.Println(FunctionGreeting)
}

// Function with Parameters
func greet(name string) string {
	fmt.Println("Hello,", name)
	return name
}

// func (u User) greet() string {
// 	return "Hello, " + u.Name
// }

// Function with Return Value
func add(a, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func PrintAny[T any](value T) {
	fmt.Println(value)
}

func checkType(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}
