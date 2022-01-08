package main

import "fmt"

// interface factory
type Person interface {
	Greet()
}

type person struct {
	Name   string
	Gender string
	Email  string
	Umur   int
}

func (p person) Greet() {
	fmt.Printf("Hi, My name is %s\n", p.Name)
}

// implementation
func NewPerson(name, gender, email string, umur int) Person {
	return person{
		Name:   name,
		Gender: gender,
		Email:  email,
		Umur:   umur,
	}
}

func main() {
	NewPerson("sahril mahendra", "male", "sahrilmahendra@example.com", 18).Greet()
	NewPerson("kanna hashimoto", "female", "hashimotokanna@example.com", 18).Greet()
}
