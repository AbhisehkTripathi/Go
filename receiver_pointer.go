package main

import "fmt"

// Person struct for demonstrating pointer receiver methods
type Person struct {
	Name string
	Age  int
}

// Pointer receiver method to update age
func (p *Person) SetAge(age int) {
	p.Age = age
}

// Pointer receiver method to update name
func (p *Person) SetName(name string) {
	p.Name = name
}

// Value receiver method (does not modify original)
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Demo function to show usage
func ReceiverPointerDemo() {
	person := Person{Name: "Abhishek", Age: 24}
	person.Greet()
	person.SetAge(30)
	person.SetName("Mohit")
	person.Greet()
}
